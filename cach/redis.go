package cach

import (
	"encoding/json"
	"film-rest-api/db"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Client *redis.Client
}

func (cache *RedisCache) SetRedis(key string, value *db.Film, expiration time.Duration) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	cache.Client.Set(key, json, expiration*time.Second)

}
func (cache *RedisCache) GetRedis(key string) *db.Film {
	val, err := cache.Client.Get(key).Result()
	if err != nil {
		return nil
	}
	film := db.Film{}
	err = json.Unmarshal([]byte(val), &film)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("get cache from redis")
	return &film
}
