package dao

import (
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	statuses struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatuses(db *sqlx.DB) repository.Statuses {
	return &statuses{db: db}
}

func(r *statuses) CreateStatus (statuses object.Statuses) (error) {
	tx, _ := r.db.Begin()
	// トランザクションの接続で操作を実行
	if _, err := tx.Exec(`insert into statuses (account_id, content) values (?, ?)`, statuses.AccountId, statuses.Content); err != nil {
		tx.Rollback()
		// 失敗だったら終了へ
		return fmt.Errorf("インサート失敗：%w", err)
	}

	if err := tx.Commit(); err != nil {
		// 失敗したら終了処理へ
		return fmt.Errorf("コミット失敗：%w", err)
	}


	return nil
}