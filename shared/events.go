package shared

import (
	"encoding/json"
	"time"
)

type OrderEvent struct {
	Event     string    `json:"event"`
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
	Amount    int       `json:"amount,omitempty"`
	Status    string    `json:"status,omitempty"`
}

func MarshalEvent(e OrderEvent) string {
	b, _ := json.Marshal(e)
	return string(b)
}

func UnmarshalEvent(payload string) (OrderEvent, error) {
	var e OrderEvent
	err := json.Unmarshal([]byte(payload), &e)
	return e, err
}
