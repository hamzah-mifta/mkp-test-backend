package domain

import (
	"context"
	"time"
)

type Theater struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TheaterRepository interface {
	IsExist(ctx context.Context, key string, val interface{}) (bool, error)
}
