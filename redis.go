package main

import (
	"strings"
	"time"

	"gopkg.in/redis.v5"
)

func init() {
	Redis = UrlClient{redis.NewClient(&redis.Options{
		Addr:     Config.RedisAddr,
		Password: Config.RedisPasswd,
		DB:       Config.RedisDB,
	})}
}

type UrlClient struct {
	*redis.Client
}

func (u *UrlClient) Set(token, path, url string, ttl time.Duration) {
	u.Client.Set(token, strings.Join([]string{path, url}, ";"), ttl)
}

func (u *UrlClient) Get(token string) (path, url string, err error) {
	v, err := u.Client.Get(token).Result()
	if err != nil {
		return "", "", err
	}

	s := strings.Split(v, ";")

	return s[0], s[1], nil
}

var Redis UrlClient
