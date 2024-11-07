package request

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateScheduleReq struct {
	MovieID   int64     `json:"movie_id"`
	TheaterID int64     `json:"theater_id"`
	ShowDate  time.Time `json:"show_date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (request CreateScheduleReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.MovieID, validation.Required),
		validation.Field(&request.TheaterID, validation.Required),
		validation.Field(&request.ShowDate, validation.Required),
		validation.Field(&request.StartTime, validation.Required),
		validation.Field(&request.EndTime, validation.Required),
	)
}

type UpdateScheduleReq struct {
	ShowDate  time.Time `json:"show_date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (request UpdateScheduleReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.ShowDate, validation.Required),
		validation.Field(&request.StartTime, validation.Required),
		validation.Field(&request.EndTime, validation.Required),
	)
}
