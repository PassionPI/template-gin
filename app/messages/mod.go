package messages

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"

	"app-ink/app/core"
)

type Messages struct {
	core         *core.Core
	limiter      *rate.Limiter
	groupName    string
	consumerName string
	streamNames  []string
	messageChan  chan MessageChan
	waitGroup    sync.WaitGroup
}

func New(core *core.Core) *Messages {
	return &Messages{
		core:         core,
		limiter:      rate.NewLimiter(1, 1),
		groupName:    "my-group",
		consumerName: "my-consumer",
		streamNames:  []string{"stream-hi", "stream-hello"},
		messageChan:  make(chan MessageChan),
		waitGroup:    sync.WaitGroup{},
	}
}

type MessageChan struct {
	Tag     string
	Stream  string
	Message redis.XMessage
}

func (m *Messages) Listen(ctx context.Context) {
	m.createGroup(ctx)
	m.readPending(ctx)
	m.readStreaming(ctx)
	m.handleMessages(ctx)
}

func (m *Messages) createGroup(
	ctx context.Context,
) {
	rds := m.core.Dep.Rds.Client
	for _, streamName := range m.streamNames {
		// 尝试创建消费者组，如果已经存在则忽略错误
		err := rds.XGroupCreateMkStream(ctx, streamName, m.groupName, "0").Err()
		if err != nil && err != redis.Nil &&
			err.Error() != "BUSYGROUP Consumer Group name already exists" {
			log.Fatalf("Could not create group: %v", err)
		}
	}
}

func (m *Messages) readPending(
	ctx context.Context,
) {
	rds := m.core.Dep.Rds.Client
	for _, streamName := range m.streamNames {
		m.waitGroup.Add(1)
		go func(streamName string) {
			defer m.waitGroup.Done()
			for {
				// 等待下一个可用的令牌
				err := m.limiter.Wait(ctx)
				if err != nil {
					log.Fatalf("Rate limiter error: %v", err)
				}
				// 读取未处理的消息
				pending, err := rds.XPending(ctx, streamName, m.groupName).Result()

				if err != nil {
					log.Fatalf("Could not get pending messages: %v", err)
				}

				if pending.Count > 0 {
					// 获取未处理的消息
					pendingMessages, err := rds.XPendingExt(ctx, &redis.XPendingExtArgs{
						Stream:   streamName,
						Group:    m.groupName,
						Start:    pending.Lower,
						End:      pending.Higher,
						Count:    1,
						Consumer: m.consumerName,
					}).Result()
					if err != nil {
						log.Fatalf("Could not get pending messages details: %v", err)
					}

					for _, pendingMessage := range pendingMessages {
						// 将未处理的消息分配给当前消费者
						claimedMessages, err := rds.XClaim(ctx, &redis.XClaimArgs{
							Stream:   streamName,
							Group:    m.groupName,
							Consumer: m.consumerName,
							MinIdle:  0,
							Messages: []string{pendingMessage.ID},
						}).Result()
						if err != nil {
							log.Fatalf("Could not claim message: %v", err)
						}

						for _, message := range claimedMessages {
							m.messageChan <- MessageChan{
								Tag:     "pending",
								Stream:  streamName,
								Message: message,
							}
						}
					}
				} else {
					break
				}
			}
		}(streamName)
	}
}

func (m *Messages) readStreaming(
	ctx context.Context,
) {
	rds := m.core.Dep.Rds.Client
	for _, streamName := range m.streamNames {
		go func(streamName string) {
			m.waitGroup.Wait()
			for {
				// 等待下一个可用的令牌
				err := m.limiter.Wait(ctx)
				if err != nil {
					log.Fatalf("Rate limiter error: %v", err)
				}
				// 从消费者组读取未处理的消息
				streams, err := rds.XReadGroup(ctx, &redis.XReadGroupArgs{
					Group:    m.groupName,
					Consumer: m.consumerName,
					Streams:  []string{streamName, ">"},
					Count:    1,
					Block:    0,
				}).Result()

				if err != nil {
					log.Fatalf("Could not read from stream: %v", err)
				}

				for _, stream := range streams {
					for _, message := range stream.Messages {
						m.messageChan <- MessageChan{
							Stream:  stream.Stream,
							Message: message,
						}
					}
				}
			}
		}(streamName)
	}
}

func (m *Messages) handleMessages(
	ctx context.Context,
) {
	rds := m.core.Dep.Rds.Client
	for {
		select {
		case ch := <-m.messageChan:
			values := ch.Message.Values
			stream := ch.Stream

			log.Printf("Processing message from %s-%s: %v", ch.Tag, stream, values["message"])

			_, err := rds.XAck(ctx, stream, m.groupName, ch.Message.ID).Result()
			if err != nil {
				log.Fatalf("Could not acknowledge message: %v", err)
			}

			_, err = rds.XDel(ctx, stream, ch.Message.ID).Result()
			if err != nil {
				log.Fatalf("Could not delete message: %v", err)
			}
		}
	}
}
