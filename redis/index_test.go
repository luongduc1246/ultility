package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestConnectRedis(t *testing.T) {
	ConnectClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Redis@Natural1246",
	})
}

func TestPut(t *testing.T) {
	r := ConnectClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Redis@Natural1246",
	})
	r.Set(context.Background(), "zing", "babe", 10*time.Second)
	val, err := r.Get(context.Background(), "zing").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func TestGet(t *testing.T) {
	r := ConnectClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Redis@Natural1246",
	})
	val, err := r.Get(context.Background(), "bDAuJBGLvRjZDfBBWCW8201692504006").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
