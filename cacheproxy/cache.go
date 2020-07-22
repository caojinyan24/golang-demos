package cacheproxy

import (
	"fmt"
	"time"
)

var CacheMap = make(map[string]Cache)

type Cache struct {
	Value      string
	ExpireTime time.Time
}

func withithCache(cacheKey string, expireTime int64, function func() (string, error)) (string, error) {
	if value, ok := CacheMap[cacheKey]; ok {
		if value.ExpireTime.After(time.Now()) {
			fmt.Println("get from cache")
			return value.Value, nil
		}
	}
	value, err := function()
	if err != nil {
		return "", err
	}
	CacheMap[cacheKey] = Cache{
		Value:      value,
		ExpireTime: time.Now().Add(time.Duration(expireTime) * time.Second),
	}
	return value, nil
}
