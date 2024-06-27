package in_memory

import (
	redis "github.com/redis/go-redis/v9"
	"rate-limiter/internal/model"
)

type RedisRepository struct {
	data        map[string]*model.Client
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *RedisRepository {
	return &RedisRepository{redisClient: redisClient}
}

func (r *RedisRepository) Upsert(client *model.Client) (*model.Client, error) {
	return nil, nil
}
func (r *RedisRepository) Create(client *model.Client) (*model.Client, error) {
	return nil, nil
}
