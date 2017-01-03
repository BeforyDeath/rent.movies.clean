package service

type MovieService interface {
	GetOne(ID int) (interface{}, error)
	GetAll(p map[string]interface{}) (interface{}, error)
}
