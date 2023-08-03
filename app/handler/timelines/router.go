package timeline

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	ar repository.Account
	sr repository.Statuses
	tr repository.TimeLine
}

func NewRouter(ar repository.Account, sr repository.Statuses, tr repository.TimeLine) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, sr, tr}
	// r.With(auth.Middleware(ar)).Get("/home", h.Home)
	r.Get("/public", h.Public)

	return r
}
