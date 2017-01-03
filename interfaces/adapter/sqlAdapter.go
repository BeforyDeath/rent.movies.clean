package adapter

type SQLAdapter interface {
	Query(query string, args ...interface{}) (SQLRows, error)
	QueryRow(query string, args ...interface{}) (SQLRow, error)
	Exec(query string, args ...interface{}) (SQLResult, error)
}

type SQLRow interface {
	Scan(dest ...interface{}) error
}

type SQLRows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type SQLResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
