package main

import (
	"fmt"
	"net/http"

	"github.com/BeforyDeath/rent.movies.clean/infrastructure"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/repository"
	"github.com/BeforyDeath/rent.movies.clean/services"
)

func main() {
	var dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" // fixme

	sqlAdapter, err := infrastructure.NewPostgres(dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer sqlAdapter.Close()

	w := services.NewWebService(
		services.WebRepository{
			Genre:    repository.NewGenreRepository(sqlAdapter),
			Movie:    repository.NewMovieRepository(sqlAdapter),
			Customer: repository.NewCustomerRepository(sqlAdapter),
			Rent:     repository.NewRentRepository(sqlAdapter),
		})

	fmt.Printf("Server started %s ...\n", ":8080")           // fixme
	fmt.Println(http.ListenAndServe(":8080", w.NewRouter())) // fixme
}
