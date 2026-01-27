package model

import "time"

type Inventory struct {
	CoffeeoffeBean int
	Milk           int
}

type Drink string

const (
	Coffee Drink = "coffee"
	Latte  Drink = "latte"
)

type Statu int

const (
	Pending Statu = iota
	Making
	Complated
)

type Order struct {
	Customer string    `json:"customer"`
	Item     Drink     `json:"item"`
	Payment  float64   `json:"payment"`
	Statu    Statu     `json:"statu"`
	CreateAt time.Time `json:"create_at"`
}
