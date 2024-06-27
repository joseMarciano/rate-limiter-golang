package internal

import (
	"rate-limiter/internal/model"
)

type IClientRepository interface {
	Upsert(client *model.Client) (*model.Client, error)
	Create(client *model.Client) (*model.Client, error)
}
