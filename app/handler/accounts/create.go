package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // r.bodyの中身を確かめてる、reqに代入
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := new(object.Account) // Account型のオブジェクト作成
	account.Username = req.Username // ユーザー名セット

	println("パスワード is:",account.CheckPassword(req.Password)) // パスワードのバリデーションチェック？

	if err := account.SetPassword(req.Password); err != nil { // パスワードセット、SetPassword内でハッシュ化の処理してる
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// panic("Must Implement Account Registration")
	
	// ここでアカウント登録の処理？
	if err := h.ar.CreateAccount(*account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil { // レスポンス内容がjsonかのチェック？
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
