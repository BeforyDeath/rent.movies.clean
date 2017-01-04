package services

import (
	"github.com/BeforyDeath/rent.movies.clean/interfaces/handler"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/repository"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/web"
)

type WebRepository struct {
	Genre    *repository.GenreRepository
	Movie    *repository.MovieRepository
	Customer *repository.CustomerRepository
	Rent     *repository.RentRepository
}

func NewWebService(repo WebRepository) *web.Handlers {
	return &web.Handlers{
		Genre:    &handler.GenreHandler{Service: GenreService{Repo: repo.Genre}},
		Movie:    &handler.MovieHandler{Service: MovieService{Repo: repo.Movie}},
		Customer: &handler.CustomerHandler{Service: CustomerService{Repo: repo.Customer}},
		Rent:     &handler.RentHandler{Service: RentService{Repo: repo.Rent}},
	}
}
