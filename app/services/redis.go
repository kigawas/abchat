package services

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(opt)
}

func AddOnlineUser(ctx context.Context, userID string) error {
	return redisClient.SAdd(ctx, "online_users", userID).Err()
}

func RemoveOnlineUser(ctx context.Context, userID string) error {
	return redisClient.SRem(ctx, "online_users", userID).Err()
}

func CheckOnlineUser(ctx context.Context, userID string) (bool, error) {
	return redisClient.SIsMember(ctx, "online_users", userID).Result()
}

func ClearOnlineUsers(ctx context.Context) error {
	return redisClient.Del(ctx, "online_users").Err()
}

func PopOfflineMessages(ctx context.Context, userID string) (string, error) {
	return redisClient.LPop(ctx, fmt.Sprintf("offline-messages:%s", userID)).Result()
}

func PushOfflineMessage(ctx context.Context, userID string, message []byte) error {
	return redisClient.RPush(ctx, fmt.Sprintf("offline-messages:%s", userID), message).Err()
}

func PublishOnlineMessage(ctx context.Context, userID string, message []byte) error {
	return redisClient.Publish(ctx, fmt.Sprintf("messages:%s", userID), message).Err()
}

func SubscribeOnlineMessages(ctx context.Context, userID string) *redis.PubSub {
	return redisClient.PSubscribe(ctx, fmt.Sprintf("messages:%s", userID))
}
