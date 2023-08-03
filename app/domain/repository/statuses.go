package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Statuses interface {
	FindByStatusId(ctx context.Context, statusId int) (*object.Statuses, error)
	CreateStatus(statuses object.Statuses) error
}
