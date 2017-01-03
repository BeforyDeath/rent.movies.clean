package domain

type Movie struct {
	ID    int
	Name  string
	Desc  string
	Year  int64  `validate:"neglect"`
	Genre string `validate:"neglect"`
}

type MovieRepository interface {
	GetOne(ID int) (Movie, error)
	GetAll(p *Pagination, f *Movie) ([]Movie, error)
	GetTotal(f *Movie) (int, error)
}
