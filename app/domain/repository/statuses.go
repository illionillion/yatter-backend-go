package repository

import "yatter-backend-go/app/domain/object"

type Statuses interface {
	CreateStatus(statuses object.Statuses) error
}
