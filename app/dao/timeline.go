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

func NewTimeLine(db *sqlx.DB) repository.TimeLine  {
	return &timeline{db: db}
}

func (r *timeline) GetPublic (ctx context.Context) (*object.TimeLine, error) {
	entity := new(object.TimeLine)
	rows, err := r.db.QueryxContext(ctx, "select statuses.id AS `id`, account.username AS `account.username` from statuses join account on statuses.account_id = account.id")
	if err != nil {
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	for rows.Next() {
		// rowsの中身を一個ずつ読み込む
		status := new(object.Statuses)
		rows.StructScan(status)
		// entityの中のStatusesに1個ずつ追加する
		entity.Statuses = append(entity.Statuses, *status)
	}

	return entity, nil
}