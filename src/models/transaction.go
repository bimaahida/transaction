package models

type Transaction struct {
	Order_ID string `json:"order_id"`
	Point    int8   `json:"point"`
	User_ID  string `json:"user_id"`
}
