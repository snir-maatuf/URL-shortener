package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define the struct StorageService to Redis client to interact with redis database
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// Limited time for cache, after 3 hours it will delete
const CacheDuration = 3 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	// init DB at redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Connecting check by PING and PONG
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err)) // STOP prog
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient

	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
