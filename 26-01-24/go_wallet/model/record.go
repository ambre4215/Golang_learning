package model

import "time"

type Category string

type TransactionType int

const (
	Expense TransactionType = iota
	Income
)

type Record struct {
	ID          int
	Amount      float64
	Type        TransactionType
	Category    Category
	Description string
	Time        time.Time
}
