package models

import "time"

type Order struct {
	Order_ID  string    `json:"order_id"`
	Point     int8      `json:"point"`
	User_ID   string    `json:"user_id"`
	Items     []Item    `json:"items"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
