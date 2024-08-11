package xStream

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

// Messages redis stream consumer
type Messages struct {
	rds           *redis.Client
	limiter       *rate.Limiter
	channel       chan channel
	waitGroup     sync.WaitGroup
	group         string
	consumer      string
	streamConfigs []StreamConfig
	handler       func(ctx context.Context, stream string, value map[string]any) error
}

// NewParams the params for create a new instance of Messages
type NewParams struct {
	RedisClient   *redis.Client
	Limiter       Limiter
	Group         string         //
	Consumer      string         //
	StreamConfigs []StreamConfig //
	Handler       func(ctx context.Context, stream string, value map[string]any) error
}

// Limiter rate limiter config item
type Limiter struct {
	Limit float64
	Burst int
}

// StreamConfig redis stream config item
type StreamConfig struct {
	Name   string
	MaxLen int64 // max message queued of stream
	Count  int64 // read count for one time
}

type channel struct {
	Tag     string
	Stream  string
	Message redis.XMessage
}

// New create a new instance of Messages
func New(params *NewParams) *Messages {
	return &Messages{
		channel:       make(chan channel),
		waitGroup:     sync.WaitGroup{},
		rds:           params.RedisClient,
		limiter:       rate.NewLimiter(rate.Limit(params.Limiter.Limit), params.Limiter.Burst),
		group:         params.Group,
		consumer:      params.Consumer,
		streamConfigs: params.StreamConfigs,
		handler:       params.Handler,
	}
}

// Sender send message to stream
func Sender(
	ctx context.Context,
	RedisClient *redis.Client,
	config StreamConfig,
	message map[string]any,
) error {
	err := RedisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: config.Name,
		Values: message, //!!! Values only support <map[string]any> !!!
		MaxLen: config.MaxLen,
	}).Err()

	if err != nil {
		log.Printf("Could not add to stream: %v", err)
		return err
	}
	return nil
}

// Listen start listen stream
func (m *Messages) Listen(ctx context.Context) {
	m.createGroup(ctx)
	m.startPending(ctx)
	m.startStreaming(ctx)
	m.listener(ctx)
}

func (m *Messages) listener(
	ctx context.Context,
) {
	for {
		select {
		case ch := <-m.channel:
			values := ch.Message.Values
			stream := ch.Stream

			m.handler(ctx, stream, values)

			_, err := m.rds.XAck(ctx, stream, m.group, ch.Message.ID).Result()
			if err != nil {
				log.Printf("Could not acknowledge message: %v", err)
			}

			_, err = m.rds.XDel(ctx, stream, ch.Message.ID).Result()
			if err != nil {
				log.Printf("Could not delete message: %v", err)
			}
		}
	}
}

func (m *Messages) createGroup(
	ctx context.Context,
) {
	for _, config := range m.streamConfigs {
		// 尝试创建消费者组，如果已经存在则忽略错误
		err := m.rds.XGroupCreateMkStream(ctx, config.Name, m.group, "0").Err()
		if err != nil && err != redis.Nil &&
			err.Error() != "BUSYGROUP Consumer Group name already exists" {
			log.Printf("Could not create group: %v", err)
		}
	}
}

func (m *Messages) startPending(
	ctx context.Context,
) {
	for _, config := range m.streamConfigs {
		m.waitGroup.Add(1)
		go func(config StreamConfig) {
			defer m.waitGroup.Done()
			m.readPending(ctx, config)
		}(config)
	}
}

func (m *Messages) readPending(
	ctx context.Context,
	config StreamConfig,
) {
	for {
		// 等待下一个可用的令牌
		err := m.limiter.Wait(ctx)
		if err != nil {
			log.Printf("Rate limiter error: %v", err)
		}
		// 读取未处理的消息
		pending, err := m.rds.XPending(ctx, config.Name, m.group).Result()

		if err != nil {
			log.Printf("Could not get pending messages: %v", err)
		}

		if pending.Count > 0 {
			// 获取未处理的消息
			pendingMessages, err := m.rds.XPendingExt(ctx, &redis.XPendingExtArgs{
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
				claimed, err := m.rds.XClaim(ctx, &redis.XClaimArgs{
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
					m.channel <- channel{
						Tag:     "XClaim",
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
	for _, config := range m.streamConfigs {
		go func(config StreamConfig) {
			m.waitGroup.Wait()
			m.readStreaming(ctx, config)
		}(config)
	}
}

func (m *Messages) readStreaming(
	ctx context.Context,
	config StreamConfig,
) {
	for {
		// 等待下一个可用的令牌
		err := m.limiter.Wait(ctx)
		if err != nil {
			log.Printf("Rate limiter error: %v", err)
		}
		// 从消费者组读取未处理的消息
		streams, err := m.rds.XReadGroup(ctx, &redis.XReadGroupArgs{
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
				m.channel <- channel{
					Stream:  stream.Stream,
					Message: message,
				}
			}
		}
	}
}
