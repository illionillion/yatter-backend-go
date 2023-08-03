package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type TimeLine interface {
	GetPublic(ctx context.Context) (*object.TimeLine, error)
}
