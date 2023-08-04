package timeline

import (
	"net/http"
	"strconv"
)
/*
GETのクエリパラメータを返す
*/
func GetParams(r *http.Request) (only_media bool, max_id int, since_id int, limit int) {
	// パラメータ取得、ハンドリング
	only_media = false
	max_id = 0
	since_id = 0
	limit = 0

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

	return only_media, max_id, since_id, limit
}
