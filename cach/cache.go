package cach

import (
	"film-rest-api/db"
	"time"
)

type CacheRedis interface {
	SetRedis(key string, value *db.Film, expiration time.Duration)
	GetRedis(key string) *db.Film
}

type CacheApp interface {
	SetApp(key string, data interface{}, expiration time.Duration) error
	GetApp(key string) ([]byte, error)
}
