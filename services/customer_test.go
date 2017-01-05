package services

import (
	"testing"

	"github.com/BeforyDeath/rent.movies.clean/domain"
)

type CR struct{}

func (repo CR) Create(c *domain.Customer) error { return nil }
func (repo CR) Login(c *domain.Customer) error  { return nil }

func TestCustomerService_Create(t *testing.T) {
	CustomerService := new(CustomerService)
	CustomerService.Repo = CR{}

	var p = map[string]interface{}{"login": "admin", "pass": "pass"}
	err := CustomerService.Create(p)
	if err != nil {
		t.Error(err)
	}
}

func TestCustomerService_Login(t *testing.T) {
	CustomerService := new(CustomerService)
	CustomerService.Repo = CR{}

	var p = map[string]interface{}{"login": "admin", "pass": "pass"}
	token, err := CustomerService.Login(p)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Expected token")
	}
}

func TestCustomerService_CheckToken(t *testing.T) {
	CustomerService := new(CustomerService)
	CustomerService.Repo = CR{}

	p := domain.Customer{Login: "admin", Pass: "pass"}
	token, err := CustomerService.createToken(&p, "salt", 5)
	if err != nil {
		t.Error(err)
	}

	claims, err := CustomerService.CheckToken(token, "salt")
	if err != nil {
		t.Error(err)
	}

	if claims["login"] != "admin" {
		t.Error("Expected admin, got", claims["login"])
	}

	_, err = CustomerService.CheckToken("CJ9.eyJOYW1lIjoiIiw", "salt")
	if err.Error() != "token contains an invalid number of segments" {
		t.Error(err)
	}

	_, err = CustomerService.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiIiwiZXhwIjoxNDgzNjQ2MzQ5LCJsb2dpbiI6ImFkbWluIiwidXNlcklEIjowfQ.SfxL5X1a5lSaCjnY9gFrm3y1G_Qpftjv5TvsjjYil22", "salt")
	if err.Error() != "signature is invalid" {
		t.Error(err)
	}

	_, err = CustomerService.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiY2hvbyIsImV4cCI6MTQ4MzM5MDA0OSwibG9naW4iOiIxMjIyMzQxcyIsInVzZXJJRCI6OX0.z9uJWyS2S1LkwJ84nyBcaA1NqjQPglTKhlHBUPmKUAk", "salt")
	if err.Error() != "Token is expired" {
		t.Error(err)
	}
}
