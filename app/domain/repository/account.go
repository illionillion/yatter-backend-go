package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// TODO: Add Other APIs
	FindByAccountId(ctx context.Context, accountId int) (*object.Account, error)
	CreateAccount(account object.Account) error
}
