package resources

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewAPI() *API {
	redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &API{
		Cache: rdb,
	}
}
