package dao

import (
	"context"
	"database/sql"
	"errors"
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

// FindByStatusId : 投稿IDから投稿を取得
func (r *statuses) FindByStatusId(ctx context.Context, statusId int) (*object.Statuses, error) {
	entity := new(object.Statuses)
	err := r.db.QueryRowxContext(ctx, "select * from statuses where id = ?", statusId).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
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