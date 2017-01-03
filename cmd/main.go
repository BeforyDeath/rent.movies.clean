package main

import (
	"fmt"
	"net/http"

	"github.com/BeforyDeath/rent.movies.clear/infrastructure"
	"github.com/BeforyDeath/rent.movies.clear/interfaces"
	"github.com/BeforyDeath/rent.movies.clear/interfaces/handler"
	"github.com/BeforyDeath/rent.movies.clear/interfaces/repository"
	"github.com/BeforyDeath/rent.movies.clear/services"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	sqlAdapter, err := infrastructure.NewPostgres(dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer sqlAdapter.Close()

	gRepo := repository.NewGenreRepository(sqlAdapter)
	gService := services.GenreService{Repo: gRepo}
	gHandler := handler.GenreHandler{Service: gService}

	mRepo := repository.NewMovieRepository(sqlAdapter)
	mService := services.MovieService{Repo: mRepo}
	mHandler := handler.MovieHandler{Service: mService}

	cRepo := repository.NewCustomerRepository(sqlAdapter)
	cService := services.CustomerService{Repo: cRepo}
	cHandler := handler.CustomerHandler{Service: cService}

	rRepo := repository.NewRentRepository(sqlAdapter)
	rService := services.RentService{Repo: rRepo}
	rHandler := handler.RentHandler{Service: rService}

	aliceService := interfaces.NewAlice()

	router := httprouter.New()

	router.POST("/customer/create", aliceService.AddChain(cHandler.Create))
	router.POST("/customer/login", aliceService.AddChain(cHandler.Login))

	router.POST("/rent/take", aliceService.AddChain(rHandler.Take, cHandler.Authorization))
	router.POST("/rent/complete", aliceService.AddChain(rHandler.Complete, cHandler.Authorization))
	router.POST("/rent/leased", aliceService.AddChain(rHandler.GetAll, cHandler.Authorization))

	router.POST("/genre", aliceService.AddChain(gHandler.GetAll))
	router.POST("/genre/:ID", aliceService.AddChain(gHandler.GetOne))
	router.POST("/movie", aliceService.AddChain(mHandler.GetAll))
	router.POST("/movie/:ID", aliceService.AddChain(mHandler.GetOne))

	fmt.Printf("Server started %s ...\n", ":8080")
	fmt.Println(http.ListenAndServe(":8080", router))

}
