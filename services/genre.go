package services

import (
	"github.com/BeforyDeath/rent.movies.clear/domain"
	"github.com/BeforyDeath/rent.movies.clear/interfaces"
)

type GenreService struct {
	Repo domain.GenreRepository
}

func (s GenreService) GetOne(ID int) (interface{}, error) {
	res, err := s.Repo.GetOne(ID) // fixme returns
	return res, err
}

func (s GenreService) GetAll(p map[string]interface{}) (interface{}, error) {
	pages, err := interfaces.NewPagination(p)
	if err != nil {
		return nil, err
	}
	rows, err := s.Repo.GetAll(pages)

	res := struct {
		Rows interface{}
	}{
		Rows: rows,
	}
	return res, err
}
