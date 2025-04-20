package main

import (
	"fmt"
	"time"

	"github.com/snykk/simple-redis-pub-sub-go/shared"
)

func main() {
	rdb := shared.NewRedisClient()
	sub := rdb.Subscribe(shared.Ctx, "order.created")
	ch := sub.Channel()

	fmt.Println("Payment Service listening...")

	for msg := range ch {
		order, _ := shared.UnmarshalEvent(msg.Payload)

		// Simulasi pembayaran
		fmt.Println("Processing payment:", order.OrderID)
		time.Sleep(2 * time.Second)

		order.Event = "payment.success"
		order.Status = "paid"
		order.Timestamp = time.Now()

		rdb.Publish(shared.Ctx, "payment.success", shared.MarshalEvent(order))
		fmt.Println("Publish payment.success")
	}
}
