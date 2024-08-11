package messages

import (
	"app-ink/app/core"
	"app-ink/pkg/xStream"
	"context"
	"fmt"
)

type Messages struct {
	Stream *xStream.Messages
}

func New(core *core.Core) *Messages {
	limiter := xStream.Limiter{
		Limit: 1,
		Burst: 1,
	}
	receiver := &Receiver{
		core: core,
	}

	return &Messages{
		Stream: xStream.New(
			&xStream.NewParams{
				Group:         "my-group",
				Consumer:      "my-consumer",
				Handler:       receiver.handler,
				Limiter:       limiter,
				RedisClient:   core.Dep.Rds.Client,
				StreamConfigs: core.Sender.StreamConfigs,
			},
		),
	}
}

type Receiver struct {
	core *core.Core
}

func (r *Receiver) handler(ctx context.Context, stream string, value map[string]any) error {
	switch stream {
	case r.core.Sender.StreamConfigMap.StreamHi.Name:
		return r.receiveHi(ctx, value)
	default:
		return nil
	}
}

func (r *Receiver) receiveHi(
	ctx context.Context,
	value map[string]any,
) (err error) {
	fmt.Println("Processing message from hi stream: ", value["message"])
	if err != nil {
		return err
	}
	return nil
}
