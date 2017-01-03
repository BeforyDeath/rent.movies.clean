package service

type GenreService interface {
	GetOne(ID int) (interface{}, error)
	GetAll(p map[string]interface{}) (interface{}, error)
}
