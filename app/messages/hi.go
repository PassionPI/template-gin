package messages

import (
	"context"
	"fmt"
)

type SayHi struct {
	message string
}

func (m *Messages) SendHi(ctx context.Context, message SayHi) error {
	return m.send(ctx, m.StreamConfigs.StreamHi, message)
}

func (m *Messages) receiveHi(
	ctx context.Context,
	value map[string]any,
) (err error) {
	fmt.Println("Processing message from hi stream: ", value["message"])
	if err != nil {
		return err
	}
	return nil
}
