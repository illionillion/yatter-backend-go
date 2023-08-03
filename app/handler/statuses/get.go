package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"

	"github.com/go-chi/chi/v5"
)

// /v1/statuses/[id]
func (h *handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusId, err := strconv.Atoi(chi.URLParam(r, "statusID"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if status, err := h.sr.FindByStatusId(ctx, statusId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if status == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else {

		if account, err := h.ar.FindByAccountId(ctx, int(status.AccountId)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if account == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else {

			entity := new(object.Statuses)
			entity.ID = status.ID
			entity.Content = status.Content
			entity.CreateAt = status.CreateAt
			entity.Account = *account

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(entity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
