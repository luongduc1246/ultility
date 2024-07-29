package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var clientRing *redis.Ring

func ConnectRing(otp *redis.RingOptions) *redis.Ring {
	if clientRing != nil {
		return clientRing
	}
	clientRing = redis.NewRing(
		otp,
	)
	err := clientRing.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	return clientRing
}
