package repository

import (
	"strings"

	"github.com/BeforyDeath/rent.movies.clean/domain"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/adapter"
)

type MovieRepository struct {
	sql adapter.SQLAdapter
}

func NewMovieRepository(a adapter.SQLAdapter) *MovieRepository {
	return &MovieRepository{sql: a}
}

func (repo MovieRepository) GetOne(ID int) (domain.Movie, error) {
	movie := domain.Movie{}
	row, err := repo.sql.QueryRow(sql["movieOne"], ID) // fixme
	if err != nil {
		return movie, err
	}
	err = row.Scan(&movie.ID, &movie.Name)
	return movie, err
}

func (repo MovieRepository) GetAll(p *domain.Pagination, f *domain.Movie) ([]domain.Movie, error) {
	repo.queryConstruct(f)
	rows, err := repo.query(p, f)
	defer rows.Close() // todo return err
	if err != nil {
		return nil, err
	}

	result := make([]domain.Movie, 0)
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Year, &movie.Genre, &movie.Desc)
		if err != nil {
			return nil, err
		}
		result = append(result, movie)
	}
	return result, nil
}

func (repo MovieRepository) GetTotal(f *domain.Movie) (int, error) {
	repo.queryConstruct(f)
	tc, err := repo.queryTotal(f)
	return tc, err
}

func (repo MovieRepository) queryConstruct(f *domain.Movie) {
	var sqlBuilt, sqlBuiltCount string
	buildType := "limit"

	sqlBuilt += sql[`movieSelect`]
	sqlBuiltCount += sql[`movieCount`]

	sqlBuilt += sql[`movieFrom`]
	sqlBuiltCount += sql[`movieFrom`]

	if f.Year > 0 {
		sqlBuilt += sql[`movieYear`]
		sqlBuiltCount += sql[`movieYear`]
		buildType += "Year"
	}
	if f.Genre != "" {
		sqlBuilt += sql[`movieGenre`]
		sqlBuiltCount += sql[`movieGenre`]
		buildType += "Genre"
	}
	sqlBuilt += sql[`movieGroup`]
	sqlBuilt += sql[`movieLimit`]

	if buildType == "limitYear" {
		sqlBuilt = strings.Replace(sqlBuilt, "$4", "$3", 1)
	}
	if buildType == "limitGenre" {
		sqlBuiltCount = strings.Replace(sqlBuiltCount, "$3", "$1", 1)
	}
	sqlBuiltCount = strings.Replace(sqlBuiltCount, "$3", "$2", 1)
	sqlBuiltCount = strings.Replace(sqlBuiltCount, "$4", "$1", 1)

	sql["movieType"] = buildType
	sql["movieBuilt"] = sqlBuilt
	sql["movieBuiltCount"] = sqlBuiltCount
}

func (repo MovieRepository) query(p *domain.Pagination, f *domain.Movie) (adapter.SQLRows, error) {
	var (
		err  error
		rows adapter.SQLRows
	)
	switch sql["movieType"] {
	case "limit":
		rows, err = repo.sql.Query(sql["movieBuilt"], p.Limit, p.Offset)
	case "limitYear":
		rows, err = repo.sql.Query(sql["movieBuilt"], p.Limit, p.Offset, f.Year)
	case "limitGenre":
		rows, err = repo.sql.Query(sql["movieBuilt"], p.Limit, p.Offset, f.Genre)
	case "limitYearGenre":
		rows, err = repo.sql.Query(sql["movieBuilt"], p.Limit, p.Offset, f.Genre, f.Year)
	}
	return rows, err
}

func (repo MovieRepository) queryTotal(f *domain.Movie) (int, error) {
	var (
		err error
		row adapter.SQLRow
		tc  int
	)
	switch sql["movieType"] {
	case "limit":
		row, err = repo.sql.QueryRow(sql["movieBuiltCount"])
	case "limitYear":
		row, err = repo.sql.QueryRow(sql["movieBuiltCount"], f.Year)
	case "limitGenre":
		row, err = repo.sql.QueryRow(sql["movieBuiltCount"], f.Genre)
	case "limitYearGenre":
		row, err = repo.sql.QueryRow(sql["movieBuiltCount"], f.Year, f.Genre)
	}
	if err != nil {
		return 0, err
	}
	err = row.Scan(&tc)
	return tc, err
}
