package dao

import (
	"context"
	"fmt"
	"strconv"
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

func (r *timeline) GetPublic(ctx context.Context, only_media bool, max_id int, since_id int, limit int) (*object.TimeLine, error) {
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

	// Add conditions for max_id and since_id
	if max_id > 0 && since_id > 0 {
		query += " WHERE statuses.id <= " + strconv.Itoa(max_id) + " AND statuses.id >= " + strconv.Itoa(since_id)
	} else if max_id > 0 {
		query += " WHERE statuses.id <= " + strconv.Itoa(max_id)
	} else if since_id > 0 {
		query += " WHERE statuses.id >= " + strconv.Itoa(since_id)
	}

	// Add limit
	if limit > 0 {
		query += " LIMIT " + strconv.Itoa(limit)
	}

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
