package database

import (
	"context"
	"fmt"
	"log"
	"video-api/pkg/config"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	/**redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "127.0.0.1"
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}**/
	redisConfig := config.Conf.Redis
	redisAddr := fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)

	RDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis:%v", err)

	}
	log.Println("Successfully connected to redis")
}
