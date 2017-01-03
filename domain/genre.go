package domain

type Genre struct {
	ID   int
	Name string
}

type GenreRepository interface {
	GetOne(ID int) (Genre, error)
	GetAll(p *Pagination) ([]Genre, error)
}
