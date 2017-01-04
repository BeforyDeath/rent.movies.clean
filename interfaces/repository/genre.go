package repository

import (
	"github.com/BeforyDeath/rent.movies.clean/domain"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/adapter"
)

type GenreRepository struct {
	sql adapter.SQLAdapter
}

func NewGenreRepository(a adapter.SQLAdapter) *GenreRepository {
	return &GenreRepository{sql: a}
}

func (repo GenreRepository) GetOne(ID int) (domain.Genre, error) {
	genre := domain.Genre{}
	row, err := repo.sql.QueryRow(sql["genreOne"], ID)
	if err != nil {
		return genre, err
	}
	err = row.Scan(&genre.ID, &genre.Name)
	return genre, err
}

func (repo GenreRepository) GetAll(p *domain.Pagination) ([]domain.Genre, error) {
	result := make([]domain.Genre, 0)
	rows, err := repo.sql.Query(sql["genreAll"], p.Limit, p.Offset)
	defer rows.Close()
	if err != nil {
		return result, err
	}
	for rows.Next() {
		genre := domain.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return result, err
		}
		result = append(result, genre)
	}
	return result, nil
}
