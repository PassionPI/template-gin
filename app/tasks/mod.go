package tasks

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	redisLock "github.com/go-co-op/gocron-redis-lock"
	"github.com/redis/go-redis/v9"
)

type Tasks struct {
	Scheduler *gocron.Scheduler
}

func New(redisClient *redis.Client) *Tasks {
	Scheduler := gocron.NewScheduler(time.UTC)
	locker, err := redisLock.NewRedisLocker(redisClient, redisLock.WithTries(1))
	if err != nil {
		panic(fmt.Errorf("failed to create redis locker: %w", err))
	}
	Scheduler.WithDistributedLocker(locker)
	return &Tasks{
		Scheduler,
	}
}
