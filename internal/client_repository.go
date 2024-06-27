package internal

import (
	"rate-limiter/internal/model"
)

type IClientRepository interface {
	Create(client *model.Client) (*model.Client, error)
}
