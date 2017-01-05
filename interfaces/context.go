package interfaces

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"     // fixme do adapter
	"github.com/julienschmidt/httprouter" // fixme do adapter
)

func GetRouterParams(r *http.Request) httprouter.Params {
	ctx := r.Context()
	return ctx.Value("params").(httprouter.Params) // fixme
}

func GetClaims(r *http.Request) jwt.MapClaims {
	ctx := r.Context()
	return ctx.Value("claims").(jwt.MapClaims) // fixme
}
