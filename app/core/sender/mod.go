package sender

import (
	"app-ink/app/core/dependency"
	"app-ink/pkg/xStream"
	"context"
)

type Sender struct {
	dep             *dependency.Dependency
	StreamConfigMap StreamConfigMap
}

type StreamConfigMap struct {
	StreamHi    xStream.StreamConfig
	StreamHello xStream.StreamConfig
}

func New(dep *dependency.Dependency) *Sender {
	return &Sender{
		dep: dep,
		StreamConfigMap: StreamConfigMap{
			StreamHi:    xStream.StreamConfig{Name: "stream-hi", MaxLen: 1000, Count: 1},
			StreamHello: xStream.StreamConfig{Name: "stream-hello", MaxLen: 1000, Count: 1},
		},
	}
}

func (s *Sender) send(ctx context.Context, config xStream.StreamConfig, message map[string]any) error {
	return xStream.Sender(ctx, s.dep.Rds.Client, config, message)
}

type SayHi struct {
	Message string
}

func (s *Sender) SendHi(ctx context.Context, message SayHi) error {
	msg := map[string]any{"message": message.Message}
	return s.send(ctx, s.StreamConfigMap.StreamHi, msg)
}
