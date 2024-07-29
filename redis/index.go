package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func ConnectClient(otp *redis.Options) *redis.Client {
	if client != nil {
		return client
	} else {
		client = redis.NewClient(
			otp,
		)
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		return client
	}
}
