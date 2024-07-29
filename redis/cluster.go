package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var clientCluster *redis.ClusterClient

func ConnectCluster(otp *redis.ClusterOptions) *redis.ClusterClient {
	if clientCluster != nil {
		return clientCluster
	}
	clientCluster = redis.NewClusterClient(
		otp,
	)
	err := clientCluster.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	return clientCluster

}
