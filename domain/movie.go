package domain

import (
	"context"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Duration  int64     `json:"duration"`
	Genre     []string  `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MovieRepository interface {
	IsExist(ctx context.Context, key string, val interface{}) (bool, error)
}
