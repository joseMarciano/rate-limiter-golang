package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"rate-limiter/application"
	"rate-limiter/infra/repository/client"
	redisRepository "rate-limiter/infra/repository/client/redis"
	"rate-limiter/internal/model"
	"strings"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error on load .env file")
	}
}

func main() {
	redisClient := redisConnection()
	redisRepo := redisRepository.NewRedisRepository(redisClient)
	defaultRepository := client.NewDefaultRepository(redisRepo)
	rateLimiterService := application.NewRateLimiterService(defaultRepository)

	engine := gin.New()
	engine.Use(limiter(rateLimiterService))
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
			"time":    time.Now(),
		})
	})

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func limiter(service *application.RateLimiterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("API_KEY")
		value := strings.Split(c.Request.RemoteAddr, ":")[0]
		typeClient := model.Ip

		if apiKey != "" {
			value = apiKey
			typeClient = model.ApiKey
		}

		if !service.Allow(value, typeClient) {
			// escrever a mensagem aqui
			c.JSON(429, gin.H{
				"message": "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
			c.Abort()
		}

	}
}

func redisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
