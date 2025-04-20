package main

import (
	"fmt"
	"time"

	"github.com/snykk/simple-redis-pub-sub-go/shared"
)

func main() {
	rdb := shared.NewRedisClient()

	order := shared.OrderEvent{
		Event:     "order.created",
		OrderID:   "ORD001",
		UserID:    "USER999",
		Amount:    300000,
		Timestamp: time.Now(),
	}

	msg := shared.MarshalEvent(order)
	err := rdb.Publish(shared.Ctx, "order.created", msg).Err()
	if err != nil {
		fmt.Println("Failed when publish order:", err)
	} else {
		fmt.Println("Order created:", msg)
	}
}
