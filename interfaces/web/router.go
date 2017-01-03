package web

import (
	"github.com/BeforyDeath/rent.movies.clear/interfaces/handler"
	"github.com/julienschmidt/httprouter" // fixme do adapter
)

type Handlers struct {
	Genre    *handler.GenreHandler
	Movie    *handler.MovieHandler
	Customer *handler.CustomerHandler
	Rent     *handler.RentHandler
}

func (web *Handlers) NewRouter() *httprouter.Router {

	alice := NewAlice()

	router := httprouter.New()

	router.POST("/customer/create", alice.AddChain(web.Customer.Create))
	router.POST("/customer/login", alice.AddChain(web.Customer.Login))

	router.POST("/rent/take", alice.AddChain(web.Rent.Take, web.Customer.Authorization))
	router.POST("/rent/complete", alice.AddChain(web.Rent.Complete, web.Customer.Authorization))
	router.POST("/rent/leased", alice.AddChain(web.Rent.GetAll, web.Customer.Authorization))

	router.POST("/genre", alice.AddChain(web.Genre.GetAll))
	router.POST("/genre/:ID", alice.AddChain(web.Genre.GetOne))

	router.POST("/movie", alice.AddChain(web.Movie.GetAll))
	router.POST("/movie/:ID", alice.AddChain(web.Movie.GetOne))

	return router
}
