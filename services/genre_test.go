package services

import (
	"testing"

	"github.com/BeforyDeath/rent.movies.clean/domain"
)

type GR struct{}

func (repo GR) GetOne(ID int) (domain.Genre, error) {
	return domain.Genre{}, nil
}

func (repo GR) GetAll(p *domain.Pagination) ([]domain.Genre, error) {
	return make([]domain.Genre, 0), nil
}

func TestGenreService_GetOne(t *testing.T) {
	GenreService := new(GenreService)
	GenreService.Repo = GR{}

	_, err := GenreService.GetOne(1)
	if err != nil {
		t.Error(err)
	}
}

func TestGenreService_GetAll(t *testing.T) {
	GenreService := new(GenreService)
	GenreService.Repo = GR{}

	var p map[string]interface{}
	_, err := GenreService.GetAll(p)
	if err != nil {
		t.Error(err)
	}
}
