package services

import (
	"time"

	"github.com/BeforyDeath/rent.movies.clear/domain"
	"github.com/BeforyDeath/rent.movies.clear/interfaces"
)

type RentService struct {
	Repo domain.RentRepository
}

func (s RentService) Take(userID int, p map[string]interface{}) error {
	rent := new(domain.Rent)
	err := interfaces.NewValidator(p, rent)
	if err != nil {
		return err
	}

	rent.UserID = userID
	rent.CreateAt = time.Now()

	err = s.Repo.Create(rent)
	return err
}

func (s RentService) Complete(userID int, p map[string]interface{}) error {
	rent := new(domain.Rent)
	err := interfaces.NewValidator(p, rent)
	if err != nil {
		return err
	}

	rent.UserID = userID
	rent.CloseAt = time.Now()

	err = s.Repo.Complete(rent)
	return err
}

func (s RentService) GetAll(userID int, p map[string]interface{}) (interface{}, error) {
	f := new(domain.RentFilter)
	err := interfaces.NewValidator(p, f)
	if err != nil {
		return nil, err
	}
	pages, err := interfaces.NewPagination(p)
	if err != nil {
		return nil, err
	}

	f.UserID = userID

	rows, err := s.Repo.GetAll(pages, f)
	if err != nil {
		return nil, err
	}

	res := struct {
		Rows interface{}
	}{
		Rows: rows,
	}
	return res, nil
}
