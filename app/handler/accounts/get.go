package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// /v1/accounts/[username]
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "userName")
	if account, err := h.ar.FindByUsername(ctx, username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if account == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(account); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}