package subscriber

import (
	"context"

	log "github.com/ebelanja/go-micro/logger"

	logspammer "github.com/ebelanja/services/logspammer/proto/logspammer"
)

type Logspammer struct{}

func (e *Logspammer) Handle(ctx context.Context, msg *logspammer.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *logspammer.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
