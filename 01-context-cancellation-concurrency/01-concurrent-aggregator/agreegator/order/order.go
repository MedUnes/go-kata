package order

import (
	"context"
	"log/slog"
	"time"
)

type Order struct {
	timeout time.Duration
}

func New(options ...func(*Order)) *Order {
	order := &Order{
		timeout: 1 * time.Second,
	}

	for _, o := range options {
		o(order)
	}
	return order
}

func WithTimeout(timeout time.Duration) func(*Order) {
	return func(o *Order) {
		o.timeout = timeout
	}
}

func (p Order) Fetch(ctx context.Context) (int, error) {
	select {
	case <-ctx.Done():
		slog.InfoContext(ctx, "recieved done signal in order service")
		return 0, ctx.Err()
	case <-time.After(p.timeout):
		slog.InfoContext(ctx, "recieved data from order service")
		return 5, nil
	}
}
