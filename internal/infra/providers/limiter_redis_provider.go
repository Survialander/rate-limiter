package infra

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type LimiterRedisProvider struct {
	client *redis.Client
}

func GetLimitterRedisProvider() *LimiterRedisProvider {

	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADDR"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	err := client.Ping(context.TODO()).Err()
	if err != nil {
		log.Fatal(err)
	}

	return &LimiterRedisProvider{
		client: client,
	}
}

func (r *LimiterRedisProvider) BlockCheck(ctx context.Context, key string) (bool, error) {
	_, err := r.client.Get(ctx, fmt.Sprintf("blocked:%v", key)).Result()

	if err != nil && err != redis.Nil {
		return false, err
	}

	if err == redis.Nil {
		return false, nil
	}

	return true, nil
}

func (r *LimiterRedisProvider) Block(ctx context.Context, key string, blockTimeInSeconds int) error {
	err := r.client.Set(ctx, fmt.Sprintf("blocked:%v", key), "", time.Second*time.Duration(blockTimeInSeconds)).Err()

	return err
}

func (r *LimiterRedisProvider) GetLimitInfo(ctx context.Context, key string) (int, error) {
	data, err := r.client.Get(ctx, key).Result()
	value, _ := strconv.Atoi(data)

	return value, err
}

func (r *LimiterRedisProvider) SetLimitInfo(ctx context.Context, key string) error {
	data, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		r.client.Set(ctx, key, 1, time.Second)
		return nil
	}

	value, _ := strconv.Atoi(data)
	r.client.Set(ctx, key, value+1, time.Second)

	return nil
}
