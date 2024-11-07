package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/hamzah-mifta/mkp-test-backend/domain"
	"github.com/hamzah-mifta/mkp-test-backend/transport/request"
	"github.com/hamzah-mifta/mkp-test-backend/utils"
)

type scheduleUsecase struct {
	scheduleRepo domain.ScheduleRepository
	movieRepo    domain.MovieRepository
	theaterRepo  domain.TheaterRepository
}

func NewScheduleUsecase(scheduleRepo domain.ScheduleRepository, movieRepo domain.MovieRepository, theaterRepo domain.TheaterRepository) *scheduleUsecase {
	return &scheduleUsecase{
		scheduleRepo: scheduleRepo,
		movieRepo:    movieRepo,
		theaterRepo:  theaterRepo,
	}
}

func (u *scheduleUsecase) Create(ctx context.Context, request *request.CreateScheduleReq) (err error) {
	isMovieExist, _ := u.movieRepo.IsExist(ctx, "id", request.MovieID)
	if !isMovieExist {
		return utils.NewBadRequestError("movie does not exists")
	}

	isTheaterExist, _ := u.theaterRepo.IsExist(ctx, "id", request.TheaterID)
	if !isTheaterExist {
		return utils.NewBadRequestError("theater does not exists")
	}

	if request.ShowDate.Before(time.Now()) {
		err = utils.NewBadRequestError("show date must be in the future")
		return
	}

	if request.StartTime.After(request.EndTime) {
		err = utils.NewBadRequestError("start time must be before end time")
	}

	err = u.scheduleRepo.Create(ctx, &domain.Schedule{
		MovieID:   request.MovieID,
		TheaterID: request.TheaterID,
		ShowDate:  request.ShowDate,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
	})

	return
}

func (u *scheduleUsecase) Fetch(ctx context.Context) (schedules []domain.Schedule, err error) {
	schedules, err = u.scheduleRepo.Fetch(ctx)
	if err != nil {
		return
	}

	return
}

func (u *scheduleUsecase) GetByID(ctx context.Context, id int64) (schedule domain.Schedule, err error) {
	schedule, err = u.scheduleRepo.GetByID(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		err = utils.NewNotFoundError("schedule not found")
		return
	}

	return
}
