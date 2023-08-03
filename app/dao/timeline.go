package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeLine(db *sqlx.DB) repository.TimeLine {
	return &timeline{db: db}
}

func (r *timeline) GetPublic(ctx context.Context) (*object.TimeLine, error) {
	entity := new(object.TimeLine)
	query := `
			select statuses.*, 
			account.username "account.username", 
			account.create_at "account.create_at",
			account.id "account.id",
			account.display_name "account.display_name",
			account.avatar "account.avatar",
			account.header "account.header",
			account.note "account.note"
			from statuses 
			join account on statuses.account_id = account.id
		`
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	for rows.Next() {
		// rowsの中身を一個ずつ読み込む
		status := new(object.Statuses)
		// rows.StructScan(status)
		if err := rows.StructScan(&status); err != nil {
			return nil, err
		}
		// entityの中のStatusesに1個ずつ追加する
		entity.Statuses = append(entity.Statuses, *status)
	}

	return entity, nil
}
