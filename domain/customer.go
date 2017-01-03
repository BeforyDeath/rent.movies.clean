package domain

import "time"

type Customer struct {
	ID       int
	Login    string `validate:"required"`
	Pass     string `validate:"required"`
	Name     string `validate:"neglect"`
	Age      int64  `validate:"neglect"`
	Phone    string `validate:"neglect"`
	CreateAt time.Time
}

type CustomerRepository interface {
	Create(c *Customer) error
	Login(c *Customer) error
}
