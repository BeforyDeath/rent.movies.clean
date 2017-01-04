package repository

import (
	"errors"
	"fmt"

	"github.com/BeforyDeath/rent.movies.clean/domain"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/adapter"
)

type CustomerRepository struct {
	sql adapter.SQLAdapter
}

func NewCustomerRepository(a adapter.SQLAdapter) *CustomerRepository {
	return &CustomerRepository{sql: a}
}

func (repo CustomerRepository) Create(c *domain.Customer) error {
	err := repo.loginIsExist(c)
	if err != nil {
		return err
	}

	row, err := repo.sql.QueryRow(sql[`customerInsert`], c.Login, c.Pass, c.Name, c.Age, c.Phone, c.CreateAt)
	if err != nil {
		return err
	}
	err = row.Scan(&c.ID)
	return err
}

func (repo CustomerRepository) Login(c *domain.Customer) error {
	row, err := repo.sql.QueryRow(sql[`customerLogin`], c.Login, c.Pass)
	if err != nil {
		return err
	}
	err = row.Scan(&c.ID, &c.Login, &c.Name, &c.Age, &c.Phone, &c.CreateAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return errors.New("Bad login or password")
		}
		return err
	}
	return nil
}

func (repo CustomerRepository) loginIsExist(c *domain.Customer) error {
	row, err := repo.sql.QueryRow(sql[`customerValid`], c.Login)
	if err != nil {
		return err
	}
	err = row.Scan(&c.ID)
	if err == nil {
		return fmt.Errorf("login %v exists", c.Login)
	}
	return nil
}
