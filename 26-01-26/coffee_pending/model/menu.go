package model

import "time"

type Inventory struct {
	coffeBean int
	milk      int
}

type Drink string

const (
	Coffee Drink = "coffee"
	Latte  Drink = "latte"
)

type Statu string

const (
	Pending   Statu = "pending"
	Making    Statu = "making"
	Complated Statu = "complated"
)

type Order struct {
	Customer string    `json:"customer"`
	Item     Drink     `json:"item"`
	Payment  float64   `json:"payment"`
	Statu    Statu     `json:"statu"`
	CreateAt time.Time `json:"create_at"`
}
