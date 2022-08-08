package cach

import (
	"encoding/json"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

type AppCache struct {
	Client *cache.Cache
}

func (r *AppCache) SetApp(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r.Client.Set(key, b, expiration)
	return nil
}

func (r *AppCache) GetApp(key string) ([]byte, error) {
	res, exist := r.Client.Get(key)
	if !exist {
		return nil, nil
	}
	resByte, ok := res.([]byte)
	if !ok {
		return nil, nil
	}
	log.Println("get cache from app")
	return resByte, nil
}
