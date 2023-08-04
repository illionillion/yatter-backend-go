package timeline

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/auth"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	account := auth.AccountOf(r)

	only_media, max_id, since_id, limit := GetParams(r)

	// SQL実行
	if statuses, err := h.tr.GetHome(ctx, int(account.ID) , only_media, max_id, since_id, limit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if statuses == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statuses.Statuses); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}