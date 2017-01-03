package handler

import (
	"net/http"
	"strconv"

	"github.com/BeforyDeath/rent.movies.clear/interfaces"
	"github.com/BeforyDeath/rent.movies.clear/interfaces/service"
)

type MovieHandler struct {
	Service service.MovieService
}

func (h MovieHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params := interfaces.GetRouterParams(r)
	ID, err := strconv.Atoi(params.ByName("ID"))
	if err != nil {
		response.Error = err.Error()
		return
	}
	result, err := h.Service.GetOne(ID)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
	response.Data = result
}

func (h MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params, err := jsonRequest(r)
	if err != nil {
		response.Error = err.Error()
		return
	}
	result, err := h.Service.GetAll(params)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
	response.Data = result
}
