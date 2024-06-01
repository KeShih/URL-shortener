package store

import (
	// "context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type StoreService struct {
	redisClient *redis.Client
}

var (
	storeService = &StoreService{}
	// ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitStore() *StoreService {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		panic(fmt.Sprintf("Failed to init Redis: %v", err))
	}

	fmt.Printf("Redis started successfully: pong message = %v\n", pong)

	storeService.redisClient = client
	return storeService
}



func SaveUrlMapping(shortUrl string, longUrl string) {
	err := storeService.redisClient.Set(shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to save URL mapping: %v", err))
	}
}

func GetLongUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to get long URL: %v", err))
	}
	return result
}


