package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/BeforyDeath/rent.movies.clean/domain"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/adapter"
)

type RentRepository struct {
	sql adapter.SQLAdapter
}

func NewRentRepository(a adapter.SQLAdapter) *RentRepository {
	return &RentRepository{sql: a}
}

func (repo RentRepository) Create(r *domain.Rent) error { // fixme add Tx
	err := repo.rentIsExist(r, false)
	if err != nil {
		return err
	}
	row, err := repo.sql.QueryRow(sql[`rentInsert`], r.UserID, r.MovieID, true, r.CreateAt)
	if err != nil {
		return err
	}
	err = row.Scan(&r.ID)
	return repo.rowIsExist(err, r.MovieID)
}

func (repo RentRepository) Complete(r *domain.Rent) error { // fixme add Tx
	err := repo.rentIsExist(r, true)
	if err != nil {
		return err
	}
	res, err := repo.sql.Exec(sql[`rentUpdate`], r.UserID, r.MovieID, false, r.CloseAt)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

// fixme ещё подумать
func (repo RentRepository) rentIsExist(r *domain.Rent, complete bool) error {
	row, err := repo.sql.QueryRow(sql[`rentValid`], r.UserID, r.MovieID, true)
	if err != nil {
		return err
	}
	err = row.Scan(&r.CreateAt)
	if err != nil && complete && err.Error() == "sql: no rows in result set" {
		return fmt.Errorf("An identifier of the movie %v is not leased", r.MovieID)
	}
	if err == nil && !complete {
		return fmt.Errorf("This movie %v you've already rented %v", r.MovieID, r.CreateAt.Format("02-01-2006 15:04"))
	}
	return nil
}

func (repo RentRepository) rowIsExist(err error, ID int64) error {
	if err != nil {
		if strings.HasSuffix(err.Error(), "\"rent_movie_id_fk\"") {
			return fmt.Errorf("An identifier of the movie %v does not exist", ID)
		}
	}
	return err
}

func (repo RentRepository) GetAll(p *domain.Pagination, f *domain.RentFilter) ([]domain.RentMovie, error) {
	rentAll := strings.Join([]string{sql[`rentSelect`], sql[`rentFrom`], sql[`rentGroup`], sql[`rentLimit`]}, "")

	res := make([]domain.RentMovie, 0)
	rows, err := repo.sql.Query(rentAll, p.Limit, p.Offset, f.UserID, !f.History)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	var CloseAt interface{}
	for rows.Next() {
		rm := domain.RentMovie{}
		err := rows.Scan(&rm.Active, &rm.CreateAt, &CloseAt, &rm.MovieID, &rm.Name, &rm.Year, &rm.Genre, &rm.Desc)
		if err != nil {
			return res, err
		}
		if t, ok := CloseAt.(time.Time); ok {
			rm.CloseAt = t
		}
		res = append(res, rm)
	}
	return res, nil
}
