package model

import "time"

type Priority int

const (
	High Priority = iota
	Middle
	Low
)

type Task struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Done     bool      `json:"done"`
	CreateAt time.Time `json:"create_at"`
	Priority Priority  `json:"priority"`
}
