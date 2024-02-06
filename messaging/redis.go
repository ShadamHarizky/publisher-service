package messaging

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisPublisher struct {
	client *redis.Client
}

func NewRedisPublisher(redisAddress string) *RedisPublisher {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	return &RedisPublisher{
		client: client,
	}
}

func (p *RedisPublisher) Close() {
	p.client.Close()
}

// Publisher ke Redis PUB/SUB
func (p *RedisPublisher) Publish(key string, message string) error {
	err := p.client.Publish(p.client.Context(), key, message).Err()
	if err != nil {
		return fmt.Errorf("Failed to publish to Redis: %v", err)
	}

	fmt.Printf("Published to Redis: %s\n", message)
	return nil
}
