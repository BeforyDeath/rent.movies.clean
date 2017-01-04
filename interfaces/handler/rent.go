package handler

import (
	"net/http"

	"github.com/BeforyDeath/rent.movies.clean/interfaces"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/service"
)

type RentHandler struct {
	Service service.RentService
}

func (h RentHandler) Take(w http.ResponseWriter, r *http.Request) {
	rw := newResponse(w)
	defer rw.Send()

	params, err := jsonRequest(r)
	if err != nil {
		rw.Error = err.Error()
		return
	}

	claims := interfaces.GetClaims(r)
	userID := int(claims["userID"].(float64))

	err = h.Service.Take(userID, params)
	if err != nil {
		rw.Error = err.Error()
		return
	}
	rw.Success = true
}

func (h RentHandler) Complete(w http.ResponseWriter, r *http.Request) {
	rw := newResponse(w)
	defer rw.Send()

	params, err := jsonRequest(r)
	if err != nil {
		rw.Error = err.Error()
		return
	}

	claims := interfaces.GetClaims(r)
	userID := int(claims["userID"].(float64))

	err = h.Service.Complete(userID, params)
	if err != nil {
		rw.Error = err.Error()
		return
	}
	rw.Success = true
}

func (h RentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rw := newResponse(w)
	defer rw.Send()

	params, err := jsonRequest(r)
	if err != nil {
		rw.Error = err.Error()
		return
	}

	claims := interfaces.GetClaims(r)
	userID := int(claims["userID"].(float64))

	result, err := h.Service.GetAll(userID, params)
	if err != nil {
		rw.Error = err.Error()
		return
	}
	rw.Success = true
	rw.Data = result
}
