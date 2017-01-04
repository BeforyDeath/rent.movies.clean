package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/BeforyDeath/rent.movies.clean/interfaces/adapter"
)

type SQLRow struct {
	Row *sql.Row
}

func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

type SQLRows struct {
	Rows *sql.Rows
}

func (r SQLRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SQLRows) Next() bool {
	return r.Rows.Next()
}

func (r SQLRows) Close() error {
	return r.Rows.Close()
}

type SQLAdapter struct {
	conn *sql.DB
}

func (a SQLAdapter) Close() error {
	if a.conn != nil {
		return a.conn.Close()
	}
	return fmt.Errorf("no open to %v", "close")
}

func (a SQLAdapter) Query(query string, args ...interface{}) (adapter.SQLRows, error) {
	err := a.checkConn()
	if err != nil {
		return SQLRows{}, err
	}
	rows, err := a.conn.Query(query, args...)
	if err != nil {
		return SQLRows{}, err
	}
	return SQLRows{Rows: rows}, nil
}

func (a SQLAdapter) QueryRow(query string, args ...interface{}) (adapter.SQLRow, error) {
	err := a.checkConn()
	if err != nil {
		return SQLRow{}, err
	}
	row := a.conn.QueryRow(query, args...)
	return SQLRow{Row: row}, nil
}

func (a SQLAdapter) Exec(query string, args ...interface{}) (adapter.SQLResult, error) {
	err := a.checkConn()
	if err != nil {
		return nil, err
	}
	return a.conn.Exec(query, args...)
}

func (a SQLAdapter) checkConn() error {
	if a.conn == nil {
		return fmt.Errorf("No Connect %v", "db")
	}
	return nil
}
