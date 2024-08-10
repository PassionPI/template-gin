package messages

import (
	"context"
	"log"
	"reflect"
	"sync"

	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"

	"app-ink/app/core"
)

type Messages struct {
	core          *core.Core
	limiter       *rate.Limiter
	channel       chan messageChannel
	waitGroup     sync.WaitGroup
	group         string
	consumer      string
	configs       []streamConfig
	StreamConfigs streamConfigs
}

type messageChannel struct {
	Tag     string
	Stream  string
	Message redis.XMessage
}

type streamConfig struct {
	Name   string
	MaxLen int64 // max message queued of stream
	Count  int64 // read count for one time
}

type streamConfigs struct {
	StreamHi    streamConfig
	StreamHello streamConfig
}

func New(core *core.Core) *Messages {
	StreamConfigs := streamConfigs{
		StreamHi:    streamConfig{Name: "stream-hi", MaxLen: 1000, Count: 1},
		StreamHello: streamConfig{Name: "stream-hello", MaxLen: 1000, Count: 1},
	}

	var configs []streamConfig
	v := reflect.ValueOf(StreamConfigs)
	for i := 0; i < v.NumField(); i++ {
		if field, ok := v.Field(i).Interface().(streamConfig); ok {
			configs = append(configs, field)
		}
	}

	return &Messages{
		core:          core,
		limiter:       rate.NewLimiter(1, 1),
		channel:       make(chan messageChannel),
		waitGroup:     sync.WaitGroup{},
		group:         "my-group",
		consumer:      "my-consumer",
		configs:       configs,
		StreamConfigs: StreamConfigs,
	}
}

func (m *Messages) send(
	ctx context.Context,
	config streamConfig,
	message any,
) error {
	err := m.core.Dep.Rds.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: config.Name,
		Values: message,
		MaxLen: config.MaxLen,
	}).Err()

	if err != nil {
		log.Printf("Could not add to stream: %v", err)
		return err
	}
	return nil
}

func (m *Messages) listener(
	ctx context.Context,
) {
	rds := m.core.Dep.Rds.Client
	for {
		select {
		case ch := <-m.channel:
			values := ch.Message.Values
			stream := ch.Stream

			switch stream {
			case m.StreamConfigs.StreamHi.Name:
				m.receiveHi(ctx, values)
				// case m.StreamConfigs.StreamHello.Name:
				// 	m.receiveHello(ctx, values)
			}

			_, err := rds.XAck(ctx, stream, m.group, ch.Message.ID).Result()
			if err != nil {
				log.Printf("Could not acknowledge message: %v", err)
			}

			_, err = rds.XDel(ctx, stream, ch.Message.ID).Result()
			if err != nil {
				log.Printf("Could not delete message: %v", err)
			}
		}
	}
}

func (m *Messages) Listen(ctx context.Context) {
	m.createGroup(ctx)
	m.startPending(ctx)
	m.startStreaming(ctx)
	m.listener(ctx)
}

func (m *Messages) createGroup(
	ctx context.Context,
) {
	rds := m.core.Dep.Rds.Client
	for _, config := range m.configs {
		// 尝试创建消费者组，如果已经存在则忽略错误
		err := rds.XGroupCreateMkStream(ctx, config.Name, m.group, "0").Err()
		if err != nil && err != redis.Nil &&
			err.Error() != "BUSYGROUP Consumer Group name already exists" {
			log.Printf("Could not create group: %v", err)
		}
	}
}

func (m *Messages) startPending(
	ctx context.Context,
) {
	for _, config := range m.configs {
		m.waitGroup.Add(1)
		go func(config streamConfig) {
			defer m.waitGroup.Done()
			m.readPending(ctx, config)
		}(config)
	}
}

func (m *Messages) readPending(
	ctx context.Context,
	config streamConfig,
) {
	rds := m.core.Dep.Rds.Client
	for {
		// 等待下一个可用的令牌
		err := m.limiter.Wait(ctx)
		if err != nil {
			log.Printf("Rate limiter error: %v", err)
		}
		// 读取未处理的消息
		pending, err := rds.XPending(ctx, config.Name, m.group).Result()

		if err != nil {
			log.Printf("Could not get pending messages: %v", err)
		}

		if pending.Count > 0 {
			// 获取未处理的消息
			pendingMessages, err := rds.XPendingExt(ctx, &redis.XPendingExtArgs{
				Stream:   config.Name,
				Group:    m.group,
				Start:    pending.Lower,
				End:      pending.Higher,
				Count:    config.Count,
				Consumer: m.consumer,
			}).Result()
			if err != nil {
				log.Printf("Could not get pending messages details: %v", err)
			}

			for _, pendingMessage := range pendingMessages {
				// 将未处理的消息分配给当前消费者
				claimed, err := rds.XClaim(ctx, &redis.XClaimArgs{
					Stream:   config.Name,
					Group:    m.group,
					Consumer: m.consumer,
					MinIdle:  0,
					Messages: []string{pendingMessage.ID},
				}).Result()
				if err != nil {
					log.Printf("Could not claim message: %v", err)
				}

				for _, message := range claimed {
					m.channel <- messageChannel{
						Tag:     "pending",
						Stream:  config.Name,
						Message: message,
					}
				}
			}
		} else {
			break
		}
	}
}

func (m *Messages) startStreaming(
	ctx context.Context,
) {
	for _, config := range m.configs {
		go func(config streamConfig) {
			m.waitGroup.Wait()
			m.readStreaming(ctx, config)
		}(config)
	}
}

func (m *Messages) readStreaming(
	ctx context.Context,
	config streamConfig,
) {
	rds := m.core.Dep.Rds.Client
	for {
		// 等待下一个可用的令牌
		err := m.limiter.Wait(ctx)
		if err != nil {
			log.Printf("Rate limiter error: %v", err)
		}
		// 从消费者组读取未处理的消息
		streams, err := rds.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    m.group,
			Consumer: m.consumer,
			Streams:  []string{config.Name, ">"},
			Count:    config.Count,
			Block:    0,
		}).Result()

		if err != nil {
			log.Printf("Could not read from stream: %v", err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				m.channel <- messageChannel{
					Stream:  stream.Stream,
					Message: message,
				}
			}
		}
	}
}
