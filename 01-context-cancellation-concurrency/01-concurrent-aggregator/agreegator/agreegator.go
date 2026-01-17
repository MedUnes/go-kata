package agreegator

import (
	"concurrent-aggregator/agreegator/order"
	"concurrent-aggregator/agreegator/profile"
	"context"
	"log/slog"
	"time"

	"golang.org/x/sync/errgroup"
)

type UserAggregator struct {
	timeout time.Duration
	logger  slog.Logger

	Profile *profile.Profile
	Order   *order.Order
}

func New(options ...func(*UserAggregator)) *UserAggregator {
	ua := &UserAggregator{
		Profile: profile.New(),
		Order:   order.New(),
		timeout: 2 * time.Second,
	}

	for _, o := range options {
		o(ua)
	}

	return ua
}

func WithTimeout(timeout time.Duration) func(*UserAggregator) {
	return func(ua *UserAggregator) {
		ua.timeout = timeout
	}
}

func WithLogger(logger slog.Logger) func(*UserAggregator) {
	return func(ua *UserAggregator) {
		ua.logger = logger
	}
}

type UserData struct {
	Profile string
	Orders  int
}

func (ua *UserAggregator) Aggregate(ctx context.Context, id int) (UserData, error) {
	slog.InfoContext(ctx, "Start")

	ctx, cancel := context.WithTimeout(ctx, ua.timeout)
	defer cancel()

	errGrp, errCtx := errgroup.WithContext(ctx)

	var profileResult string
	var orderResult int

	errGrp.Go(func() error {
		result, err := ua.Profile.Fetch(errCtx)
		if err != nil {
			return err
		}

		profileResult = result
		return nil
	})

	errGrp.Go(func() error {
		result, err := ua.Order.Fetch(errCtx)
		if err != nil {
			return err
		}

		orderResult = result
		return nil
	})

	if err := errGrp.Wait(); err != nil {
		slog.InfoContext(errCtx, "recieved error in agreegator", "error", err)
		return UserData{}, err
	}

	slog.InfoContext(errCtx, "End")

	return UserData{
		Profile: profileResult,
		Orders:  orderResult,
	}, nil
}
