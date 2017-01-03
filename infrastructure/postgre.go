package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgres(dsn string) (*SQLAdapter, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return new(SQLAdapter), err
	}
	err = conn.Ping()
	if err != nil {
		return new(SQLAdapter), err
	}
	return &SQLAdapter{conn: conn}, nil
}
