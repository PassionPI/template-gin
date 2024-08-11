package corn

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	lock "github.com/go-co-op/gocron-redis-lock"
	"github.com/redis/go-redis/v9"
)

type Cron struct {
	Scheduler *gocron.Scheduler
}

func New(redisClient *redis.Client) *Cron {
	Scheduler := gocron.NewScheduler(time.UTC)
	locker, err := lock.NewRedisLocker(redisClient, lock.WithTries(1))
	if err != nil {
		panic(fmt.Errorf("failed to create redis locker: %w", err))
	}
	Scheduler.WithDistributedLocker(locker)
	return &Cron{
		Scheduler,
	}
}
