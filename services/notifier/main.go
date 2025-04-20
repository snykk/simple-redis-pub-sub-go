package main

import (
	"fmt"

	"github.com/snykk/simple-redis-pub-sub-go/shared"
)

func main() {
	rdb := shared.NewRedisClient()
	sub := rdb.PSubscribe(shared.Ctx, "order.*", "payment.*", "inventory.*")
	ch := sub.Channel()

	fmt.Println("Notifier Service listening...")

	for msg := range ch {
		event, _ := shared.UnmarshalEvent(msg.Payload)
		fmt.Printf("Event received [%s]: OrderID: %s, Status: %s\n",
			event.Event, event.OrderID, event.Status)
	}
}
