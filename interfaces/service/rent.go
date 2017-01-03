package service

type RentService interface {
	Take(userID int, p map[string]interface{}) error
	Complete(userID int, p map[string]interface{}) error
	GetAll(userID int, p map[string]interface{}) (interface{}, error)
}
