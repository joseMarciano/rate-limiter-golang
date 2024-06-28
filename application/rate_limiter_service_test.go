package application

import (
	"github.com/stretchr/testify/assert"
	"rate-limiter/infra/repository/client"
	"rate-limiter/internal/model"
	"testing"
	"time"
)

type RedisMock struct {
}

func (r *RedisMock) Create(client *model.Client) (*model.Client, error) {
	return client, nil
}

func TestRateLimit(t *testing.T) {
	// Initialize your rate limiter
	defaultRepository := client.NewDefaultRepository(&RedisMock{})
	rateLimiter := NewRateLimiterService(defaultRepository) // 1 request per second

	t.Run("Allow first request", func(t *testing.T) {
		allowed := rateLimiter.Allow("user1", model.Ip)
		assert.True(t, allowed, "First request should be allowed")
	})

	t.Run("Deny second request within rate limit", func(t *testing.T) {
		allowed := rateLimiter.Allow("user1", model.Ip)
		assert.False(t, allowed, "Second request within rate limit should be denied")
	})

	t.Run("Allow request after rate limit duration", func(t *testing.T) {
		time.Sleep(time.Second * 3)
		allowed := rateLimiter.Allow("user1", model.Ip)
		assert.True(t, allowed, "Request after rate limit duration should be allowed")
	})
}
