package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Rds struct {
	Client *redis.Client
}

func New(uri string) *Rds {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		fmt.Println("Redis URL parse error: ", err)
		panic(err)
	}

	Client := redis.NewClient(opt)

	return &Rds{
		Client,
	}
}

func (r *Rds) Lock(ctx context.Context, key string, duration time.Duration) bool {
	result, err := r.Client.SetNX(ctx, key, true, duration).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

func (r *Rds) Unlock(ctx context.Context, key string) bool {
	_, err := r.Client.Del(ctx, key).Result()
	if err != nil {
		return false
	}
	return true
}

func (r *Rds) IsLocked(ctx context.Context, key string) bool {
	result, err := r.Client.Exists(ctx, key).Result()
	if err != nil || result == 0 {
		return false
	}
	return true
}
