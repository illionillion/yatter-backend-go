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
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func(r *account) CreateAccount (account object.Account) (error)  {
	
	tx, _ := r.db.Begin()
	// トランザクションの接続で操作を実行
	if _, err := tx.Exec(`insert into account (username, password_hash) values (?, ?)`, account.Username, account.PasswordHash); err != nil {
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