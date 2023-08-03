package object

import "time"

type Statuses struct {
	ID int64 `json:"id,omitempty"`

	Account Account `json:"account,omitempty" db:"-"`

	AccountId int64 `json:"-" db:"account_id"`

	Content *string `json:"content"`

	CreateAt time.Time `json:"create_at,omitempty" db:"created_at"`
}