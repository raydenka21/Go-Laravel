package setting

import (
	"context"
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() *redis.Client {

	redisHost := facades.Config().Env("REDIS_HOST")
	redisPort := facades.Config().Env("REDIS_PORT")
	redisPassword := facades.Config().Env("REDIS_PASSWORD")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: fmt.Sprintf("%s", redisPassword),
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		facades.Log().Debug(err)

	}
	return rdb
}
