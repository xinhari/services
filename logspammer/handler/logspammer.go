package handler

import (
	"context"

	log "github.com/ebelanja/go-micro/logger"

	logspammer "github.com/ebelanja/micro-services/logspammer/proto/logspammer"
)

type Logspammer struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Logspammer) Call(ctx context.Context, req *logspammer.Request, rsp *logspammer.Response) error {
	log.Info("Received Logspammer.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Logspammer) Stream(ctx context.Context, req *logspammer.StreamingRequest, stream logspammer.Logspammer_StreamStream) error {
	log.Infof("Received Logspammer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&logspammer.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Logspammer) PingPong(ctx context.Context, stream logspammer.Logspammer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&logspammer.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
