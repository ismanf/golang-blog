package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type KeyNotExistsErorr struct {
	Key string
}

func (k KeyNotExistsErorr) Error() string {
	return fmt.Sprintf("Key: %s not exist in store", k.Key)
}

var client *redis.Client

func Get(key string) (string, error) {
	key, err := client.Get(key).Result()

	if err == redis.Nil {
		knErr := KeyNotExistsErorr{key}
		return key, knErr
	}

	if err != nil {
		return key, err
	}

	return key, nil
}

func Set(key string, value interface{}, exp time.Duration) error {
	val, _ := json.Marshal(value)
	return client.Set(key, val, exp).Err()
}

func Initialize() {
	client = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}
