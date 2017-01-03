package handler

import (
	"net/http"

	"github.com/BeforyDeath/rent.movies.clear/interfaces"
	"github.com/BeforyDeath/rent.movies.clear/interfaces/service"
)

type RentHandler struct {
	Service service.RentService
}

func (h RentHandler) Take(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params, err := jsonRequest(r)
	if err != nil {
		response.Error = err.Error()
		return
	}

	claims := interfaces.GetClaims(r)
	userID := int(claims["userID"].(float64))

	err = h.Service.Take(userID, params)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
}

func (h RentHandler) Complete(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params, err := jsonRequest(r)
	if err != nil {
		response.Error = err.Error()
		return
	}

	claims := interfaces.GetClaims(r)
	userID := int(claims["userID"].(float64))

	err = h.Service.Complete(userID, params)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
}

func (h RentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params, err := jsonRequest(r)
	if err != nil {
		response.Error = err.Error()
		return
	}

	claims := interfaces.GetClaims(r)
	userID := int(claims["userID"].(float64))

	result, err := h.Service.GetAll(userID, params)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
	response.Data = result
}
