package in_memory

import (
	"context"
	"encoding/json"
	redis "github.com/redis/go-redis/v9"
	"rate-limiter/internal/model"
	"time"
)

type RedisRepository struct {
	data        map[string]*model.Client
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *RedisRepository {
	return &RedisRepository{redisClient: redisClient}
}

func (r *RedisRepository) Create(client *model.Client) (*model.Client, error) {
	payload, err := json.Marshal(client)
	if err != nil {
		return nil, err
	}

	result := r.redisClient.Set(context.TODO(), client.Id, payload, time.Minute*10)

	if result.Err() != nil {
		return nil, result.Err()
	}

	return client, nil
}
