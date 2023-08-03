package timeline

import (
	"encoding/json"
	"net/http"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// パラメータ取得、ハンドリング
	// SQL実行
	if statuses, err := h.tr.GetPublic(ctx); err != nil {
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
