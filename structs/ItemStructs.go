package structs

import "time"

type Items struct {
	ItemCode int `json:"itemCode"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

type Order struct {
	OrderedAt time.Time `json:"orderedAt"`
	CustomerName string `json:"customerName"`
	Items []Items `json:"items"`
}

