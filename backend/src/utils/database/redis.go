package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Password string `json:"password"`

	Client *redis.Client
}

func New(addr string, port string, pass string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", addr, port),
		Password: pass,
		DB:       0,
	})

	rds := &Redis{addr, port, pass, client}
	return rds
}

func (rds *Redis) Ping() error {
	return rds.Client.Ping().Err()
}
