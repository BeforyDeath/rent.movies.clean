package interfaces

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type AliceService struct {
	base alice.Chain
}

func NewAlice(constructors ...alice.Constructor) *AliceService {
	as := new(AliceService)
	as.base = alice.New(constructors...)
	return as
}

func (m AliceService) AddChain(h http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	b := m.base.Extend(alice.New(c...))
	return m.stripParams(b.ThenFunc(h))
}

func (m AliceService) stripParams(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps) // fixme
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func GetRouterParams(r *http.Request) httprouter.Params {
	ctx := r.Context()
	return ctx.Value("params").(httprouter.Params) // fixme
}

func GetClaims(r *http.Request) jwt.MapClaims {
	ctx := r.Context()
	return ctx.Value("claims").(jwt.MapClaims) // fixme
}
