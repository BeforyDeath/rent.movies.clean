package services

import (
	"github.com/BeforyDeath/rent.movies.clean/domain"
	"github.com/BeforyDeath/rent.movies.clean/interfaces"
)

type MovieService struct {
	Repo domain.MovieRepository
}

func (s MovieService) GetOne(ID int) (interface{}, error) {
	res, err := s.Repo.GetOne(ID)
	return res, err
}

func (s MovieService) GetAll(p map[string]interface{}) (interface{}, error) {
	pages, err := interfaces.NewPagination(p)
	if err != nil {
		return nil, err
	}

	filter := new(domain.Movie)
	err = interfaces.NewValidator(p, filter)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.Repo.GetTotal(filter)
	if err != nil {
		return nil, err
	}

	var rows interface{}
	if totalCount > 0 {
		rows, err = s.Repo.GetAll(pages, filter)
		if err != nil {
			return nil, err
		}
	}

	res := struct {
		Rows       interface{}
		TotalCount int
	}{
		Rows:       rows,
		TotalCount: totalCount,
	}
	return res, err
}
