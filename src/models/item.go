package models

type Item struct {
	Sku      string `json:"sku"`
	Quantity int8   `json:"quantity"`
}
