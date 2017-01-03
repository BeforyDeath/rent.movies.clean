package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/BeforyDeath/rent.movies.clear/infrastructure"
)

func main() {
	var dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" // fixme

	filename := flag.String("f", "", "Initialized databases from filename")
	flag.Parse()

	if *filename != "" {

		sqlAdapter, err := infrastructure.NewPostgres(dsn)
		if err != nil {
			fmt.Println(err)
		}
		defer sqlAdapter.Close()

		path, _ := os.Getwd()
		file, err := os.Open(path + "/" + *filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f := bufio.NewReader(file)
		for {
			str, err := f.ReadString(';')
			if err != nil {
				break
			}

			_, err = sqlAdapter.Exec(str)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Printf("Initialized databases from %v\n", *filename)
	}
}
