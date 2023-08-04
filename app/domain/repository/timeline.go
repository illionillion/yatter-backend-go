package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type TimeLine interface {
	GetPublic(ctx context.Context, only_media bool, max_id int, since_id int, limit int) (*object.TimeLine, error)
}
