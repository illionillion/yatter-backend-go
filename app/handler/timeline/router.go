package timeline

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	ar repository.Account
	sr repository.Statuses
}

func NewRouter(ar repository.Account, sr repository.Statuses) http.Handler  {
	r := chi.NewRouter()

	h := &handler{ar, sr}
	// r.With(auth.Middleware(ar)).Get("/home", h.Home)
	r.Get("/public", h.Public)

	return r
}