package port

import (
	"context"
	"fmt"

	"github.com/malikbenkirane/natsio/svctl/pkg/channel"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Port struct {
	chReady chan *nats.Msg
	sub     *nats.Subscription
	log     *zap.Logger
}

func New(natsUrl string, log *zap.Logger) (*Port, error) {
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		return nil, fmt.Errorf("nats: connect %q: %w", natsUrl, err)
	}

	p := &Port{
		chReady: make(chan *nats.Msg, 64),
	}

	sub, err := nc.ChanSubscribe(channel.Ready, p.chReady)
	if err != nil {
		return nil, fmt.Errorf("subscribe %q: %w", channel.Ready, err)
	}
	p.sub = sub

	p.log = log

	return p, nil
}

func (p Port) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			if err := p.sub.Unsubscribe(); err != nil {
				return fmt.Errorf("unsubscribe: %w", err)
			}
			return ctx.Err()
		case msg := <-p.chReady:
			p.log.Info("service ready", zap.String("service_name", string(msg.Data)))
		}
	}
}
