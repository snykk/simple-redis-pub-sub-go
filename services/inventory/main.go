package main

import (
	"fmt"
	"time"

	"github.com/snykk/simple-redis-pub-sub-go/shared"
)

func main() {
	rdb := shared.NewRedisClient()
	sub := rdb.Subscribe(shared.Ctx, "payment.success")
	ch := sub.Channel()

	fmt.Println("Inventory Service listening...")

	for msg := range ch {
		payment, _ := shared.UnmarshalEvent(msg.Payload)
		fmt.Println("Reduce stock:", payment.OrderID)
		time.Sleep(1 * time.Second)

		payment.Event = "inventory.updated"
		payment.Status = "stock-adjusted"
		payment.Timestamp = time.Now()

		rdb.Publish(shared.Ctx, "inventory.updated", shared.MarshalEvent(payment))
		fmt.Println("Publish inventory.updated")
	}
}
