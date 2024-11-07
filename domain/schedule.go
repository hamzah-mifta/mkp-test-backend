package domain

import (
	"context"
	"time"

	"github.com/hamzah-mifta/mkp-test-backend/transport/request"
)

type Schedule struct {
	ID        int64     `json:"id"`
	MovieID   int64     `json:"movie_id"`
	TheaterID int64     `json:"theater_id"`
	ShowDate  time.Time `json:"show_date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ScheduleRepository interface {
	Create(ctx context.Context, schedule *Schedule) error
	GetAll(ctx context.Context) ([]Schedule, error)
	GetByID(ctx context.Context, id int64) (Schedule, error)
	Update(ctx context.Context, schedule *Schedule) error
	Delete(ctx context.Context, id int64) error
}

type ScheduleUsecase interface {
	Create(ctx context.Context, schedule *request.CreateScheduleReq) error
}
