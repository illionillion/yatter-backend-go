package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

type AddRequest struct {
	// AccountId int64
	Status string
	// Media
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

	// ヘッダーのAuthenticationからアカウント情報を取得
	account := auth.AccountOf(r)

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // r.bodyの中身を確かめてる、reqに代入
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Status投稿
	status := new(object.Statuses)
	status.AccountId = account.ID
	status.Content = &req.Status
	if err := h.sr.CreateStatus(*status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil { // レスポンス内容がjsonかのチェック？
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
