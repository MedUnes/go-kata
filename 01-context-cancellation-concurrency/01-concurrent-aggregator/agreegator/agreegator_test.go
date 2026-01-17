package agreegator_test

import (
	"concurrent-aggregator/agreegator"
	"concurrent-aggregator/agreegator/order"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAggregate(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		ag := agreegator.New()
		user, err := ag.Aggregate(context.Background(), 1)

		require.NoError(t, err)
		require.Equal(t, agreegator.UserData{Profile: "Alice", Orders: 5}, user)
	})

	t.Run("slow poke", func(t *testing.T) {
		ag := agreegator.New(agreegator.WithTimeout(1 * time.Second))
		ag.Order = order.New(order.WithTimeout(2 * time.Second))
		user, err := ag.Aggregate(context.Background(), 1)

		require.ErrorIs(t, err, context.DeadlineExceeded)
		require.Empty(t, user)
	})

	t.Run("slow poke", func(t *testing.T) {
		ag := agreegator.New(agreegator.WithTimeout(1 * time.Second))
		ag.Order = order.New(order.WithTimeout(2 * time.Second))
		user, err := ag.Aggregate(context.Background(), 1)

		require.ErrorIs(t, err, context.DeadlineExceeded)
		require.Empty(t, user)
	})
}
