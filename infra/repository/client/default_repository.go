package client

import (
	"fmt"
	"golang.org/x/time/rate"
	"rate-limiter/env"
	"rate-limiter/internal"
	"rate-limiter/internal/model"
	"time"
)

type DefaultRepository struct {
	data            map[string]*model.ClientLimiter // ajustar para ser chave e limiter somente
	redisRepository internal.IClientRepository
}

func NewDefaultRepository(redisRepository internal.IClientRepository) *DefaultRepository {
	repository := &DefaultRepository{
		data:            make(map[string]*model.ClientLimiter),
		redisRepository: redisRepository,
	}

	repository.triggerClearJob()

	return repository
}

func (r *DefaultRepository) Create(client *model.Client) *model.ClientLimiter {
	clientLimiter := &model.ClientLimiter{
		Client: client,
		Limit:  rate.NewLimiter(rate.Limit(env.GetConfigLimitRate(client.Type)), 10),
	}
	r.data[client.Id] = clientLimiter

	_, err := r.redisRepository.Create(client)
	if err != nil {
		fmt.Printf("Error on create in redis %s", err.Error())
	}

	return clientLimiter
}

func (r *DefaultRepository) Update(client *model.Client) {
	_, err := r.redisRepository.Create(client)
	if err != nil {
		fmt.Printf("Error on update in redis %s", err.Error())
	}
}

func (r *DefaultRepository) FindById(id string) *model.ClientLimiter {
	if _, ok := r.data[id]; !ok {
		return nil
	}

	return r.data[id]
}

func (r *DefaultRepository) FindAll() []*model.ClientLimiter {
	var values []*model.ClientLimiter

	for _, v := range r.data {
		values = append(values, v)
	}

	return values
}

func (r *DefaultRepository) triggerClearJob() {
	go func() {

		for {
			time.Sleep(time.Second * 10)
		}

	}()
}
