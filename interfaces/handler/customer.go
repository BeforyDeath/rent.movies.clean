package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BeforyDeath/rent.movies.clear/interfaces/service"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (h CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params, err := jsonRequest(r)
	if err != nil {
		response.Error = err.Error()
		return
	}
	err = h.Service.Create(params)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
}

func (h CustomerHandler) Login(w http.ResponseWriter, r *http.Request) {
	response := newResponse(w)
	defer response.Send()

	params, err := jsonRequest(r)
	if err != nil {
		response.Error = err.Error()
		return
	}
	result, err := h.Service.Login(params)
	if err != nil {
		response.Error = err.Error()
		return
	}
	response.Success = true
	response.Data = result
}

func (h CustomerHandler) Authorization(next http.Handler) http.Handler { // fixme response
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")
		if token != "" {
			claims, err := h.Service.CheckToken(token, "salt") // fixme
			if err != nil {

				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				re := response{}
				re.Error = err.Error()
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(re)
				return

			}
			ctx := context.WithValue(r.Context(), "claims", claims) // fixme
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		re := response{}
		re.Error = "One of the parameters specified was missing or invalid: address is token"
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(re)

		return
	}
	return http.HandlerFunc(fn)
}
