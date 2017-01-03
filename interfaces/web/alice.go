package web

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter" // fixme do adapter
	"github.com/justinas/alice"           // fixme do adapter
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
