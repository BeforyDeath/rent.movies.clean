package handler

import (
	"net/http"
	"strconv"

	"github.com/BeforyDeath/rent.movies.clean/interfaces"
	"github.com/BeforyDeath/rent.movies.clean/interfaces/service"
)

type GenreHandler struct {
	Service service.GenreService
}

func (h GenreHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	rw := newResponse(w)
	defer rw.Send()

	params := interfaces.GetRouterParams(r)
	ID, err := strconv.Atoi(params.ByName("ID"))
	if err != nil {
		rw.Error = err.Error()
		return
	}
	result, err := h.Service.GetOne(ID)
	if err != nil {
		rw.Error = err.Error()
		return
	}
	rw.Success = true
	rw.Data = result
}

func (h GenreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rw := newResponse(w)
	defer rw.Send()

	params, err := jsonRequest(r)
	if err != nil {
		rw.Error = err.Error()
		return
	}
	result, err := h.Service.GetAll(params)
	if err != nil {
		rw.Error = err.Error()
		return
	}
	rw.Success = true
	rw.Data = result
}
