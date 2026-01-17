package profile

import (
	"context"
	"log/slog"
	"time"
)

type Profile struct {
	timeout time.Duration
}

func New(options ...func(*Profile)) *Profile {
	p := &Profile{
		timeout: 1 * time.Second,
	}

	for _, o := range options {
		o(p)
	}

	return p
}

func (p Profile) Fetch(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		slog.InfoContext(ctx, "recieved done signal in profile service")
		return "", ctx.Err()
	case <-time.After(p.timeout):
		slog.InfoContext(ctx, "recieved data from profile service")
		return "Alice", nil
	}
}
