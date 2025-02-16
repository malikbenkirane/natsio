package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/malikbenkirane/natsio/svctl/internal/port"
	"go.uber.org/zap"
)

type Service struct {
	port *port.Port
}

var (
	ErrNatsUrlNotSet = errors.New("NATS_URL environment variable not set")
)

func New(log *zap.Logger) (*Service, error) {

	natsUrl := os.Getenv("NATS_URL")
	if natsUrl == "" {
		return nil, ErrNatsUrlNotSet
	}

	var (
		svc Service
		err error
	)

	svc.port, err = port.New(natsUrl, log.Named("port"))
	if err != nil {
		return nil, fmt.Errorf("new port: %w", err)
	}

	return &svc, nil

}

func (s Service) Start(ctx context.Context) error {
	return s.port.Serve(ctx)
}
