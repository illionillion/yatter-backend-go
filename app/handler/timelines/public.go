package timeline

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// パラメータ取得、ハンドリング
	only_media := false
	max_id := 0
	since_id := 0
	limit := 0

	if r.FormValue("only_media") != "" {
		val, err := strconv.ParseBool(r.FormValue("only_media"))
		if err == nil {
			only_media = val
		}
	}

	if r.FormValue("max_id") != "" {
		val, err := strconv.Atoi(r.FormValue("max_id"))
		if err == nil {
			max_id = val
		}
	}

	if r.FormValue("since_id") != "" {
		val, err := strconv.Atoi(r.FormValue("since_id"))
		if err == nil {
			since_id = val
		}
	}

	if r.FormValue("limit") != "" {
		val, err := strconv.Atoi(r.FormValue("limit"))
		if err == nil {
			limit = val
		}
	}

	// println("only_media", r.FormValue("only_media"), only_media)
	// println("max_id", r.FormValue("max_id"), max_id)
	// println("since_id", r.FormValue("since_id"), since_id)
	// println("limit", r.FormValue("limit"), limit)

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
