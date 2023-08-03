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

func (r *timeline) GetPublic (ctx context.Context, timeline object.TimeLine) (*object.TimeLine, error) {
	entity := new(object.TimeLine)
	rows, err := r.db.SelectContext(ctx, &entity, "select * from statuses")
	if err != nil {
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	// for rows.Next() {

	// }

	return nil
}