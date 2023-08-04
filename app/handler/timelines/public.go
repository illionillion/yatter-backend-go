package timeline

import (
	"encoding/json"
	"net/http"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	only_media, max_id, since_id, limit := GetParams(r)

	// SQL実行
	if statuses, err := h.tr.GetPublic(ctx, only_media, max_id, since_id, limit); err != nil {
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
