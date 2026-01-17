package main

import (
	aggregator "concurrent-aggregator/agreegator"
	"concurrent-aggregator/agreegator/order"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	ag := aggregator.New(
		aggregator.WithTimeout(5*time.Second),
		aggregator.WithLogger(*slog.Default()),
	)

	ag.Order = order.New(order.WithTimeout(10 * time.Second))

	user, err := ag.Aggregate(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
