package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/BeforyDeath/rent.movies.clear/domain"
	"github.com/BeforyDeath/rent.movies.clear/interfaces"
	"github.com/dgrijalva/jwt-go" // fixme do adapter
)

type CustomerService struct {
	Repo domain.CustomerRepository
}

func (s CustomerService) Create(p map[string]interface{}) error {
	customer := new(domain.Customer)
	err := interfaces.NewValidator(p, customer)
	if err != nil {
		return err
	}

	customer.CreateAt = time.Now()
	customer.Pass, _ = s.hashed(customer.Pass, "salt") // fixme

	err = s.Repo.Create(customer)
	return err
}

func (s CustomerService) Login(p map[string]interface{}) (string, error) {
	customer := new(domain.Customer)
	err := interfaces.NewValidator(p, customer)
	if err != nil {
		return "", err
	}

	customer.Pass, _ = s.hashed(customer.Pass, "salt") // fixme

	err = s.Repo.Login(customer)
	if err != nil {
		return "", err
	}

	token, err := s.createToken(customer, "salt") // fixme
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s CustomerService) hashed(str, salt string) (string, error) {
	mac := hmac.New(sha256.New, []byte(salt))
	_, err := mac.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}

func (s CustomerService) createToken(c *domain.Customer, salt string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["userID"] = c.ID
	claims["login"] = c.Login
	claims["Name"] = c.Name
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix() // fixme
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(salt))
	return tokenString, err
}

func (s CustomerService) CheckToken(token string, salt string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(salt), nil })
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("This token %v", "is terrible") // fixme
}
