package helper

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/redis/go-redis/v9"
	"time"
)

func RedisGet(ctx http.Context, rdb *redis.Client, key string) (string, error) {
	val2, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "empty", nil
	} else if err != nil {
		return "", err
	}
	return val2, nil
}

func RedisSet(ctx http.Context, rdb *redis.Client, key string, value []byte, time time.Duration) error {
	err := rdb.Set(ctx, key, value, time).Err()
	if err != nil {
		return err
	}
	return nil

}
