package application

import (
	"math"
	"rate-limiter/env"
	"rate-limiter/infra/repository/client"
	"rate-limiter/internal/model"
	"sync"
	"time"
)

type RateLimiterService struct {
	defaultRepository *client.DefaultRepository
	mtx               *sync.Mutex
}

func NewRateLimiterService(defaultRepository *client.DefaultRepository) *RateLimiterService {
	r := &RateLimiterService{defaultRepository: defaultRepository, mtx: &sync.Mutex{}}
	r.triggerUnlockClient()
	return r
}

func (s *RateLimiterService) Allow(id string, typeClient model.TypeClient) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	clientLimiter := s.defaultRepository.FindById(id)

	if clientLimiter == nil {
		clientLimiter = s.defaultRepository.Create(&model.Client{
			Id:       id,
			Type:     typeClient,
			LastHit:  time.Now(),
			IsLocked: false,
		})

		return clientLimiter.Limit.Allow()
	}

	if clientLimiter.Client.IsLocked {
		//fmt.Printf("Client %s is locked\r\n", clientLimiter.Client.Id)
		return false
	}

	allow := clientLimiter.Limit.Allow()

	clientLimiter.Client.LastHit = time.Now()
	if !allow {
		clientLimiter.Client.IsLocked = true
	}

	s.defaultRepository.Update(clientLimiter.Client)
	return allow
}

func (s *RateLimiterService) triggerUnlockClient() {
	go func() {
		for {
			time.Sleep(time.Second)
			s.mtx.Lock()
			//fmt.Printf("Executing unlock at %s\r\n", time.Now())
			for _, c := range s.defaultRepository.FindAll() {

				if !c.Client.IsLocked {
					continue
				}

				difference := math.Abs(time.Now().Sub(c.Client.LastHit).Seconds())
				lockedTime := float64(env.GetLockedTimeLimitRate(c.Client.Type))

				if difference >= lockedTime {
					c.Client.IsLocked = false
					s.defaultRepository.Update(c.Client)
				}

			}

			s.mtx.Unlock()
		}
	}()
}
