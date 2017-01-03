package domain

import "time"

type Rent struct {
	ID       int
	UserID   int   `json:"-"`
	MovieID  int64 `validate:"required"`
	Active   bool
	CreateAt time.Time
	CloseAt  time.Time
}

type RentMovie struct {
	Rent
	Movie
}

type RentFilter struct {
	UserID  int
	History bool `validate:"neglect"`
}

type RentRepository interface {
	Create(r *Rent) error
	Complete(r *Rent) error
	GetAll(p *Pagination, f *RentFilter) ([]RentMovie, error)
}
