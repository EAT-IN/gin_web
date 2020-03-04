package model

import (
	"time"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Status    bool   `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
