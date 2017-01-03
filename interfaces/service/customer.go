package service

import jwt "github.com/dgrijalva/jwt-go" // fixme do adapter

type CustomerService interface {
	Create(p map[string]interface{}) error
	Login(p map[string]interface{}) (string, error)
	CheckToken(token string, salt string) (jwt.MapClaims, error)
}
