package pgsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hamzah-mifta/mkp-test-backend/domain"
)

type pgsqlScheduleRepository struct {
	db *sql.DB
}

func NewPgsqlScheduleRepository(db *sql.DB) *pgsqlScheduleRepository {
	return &pgsqlScheduleRepository{
		db: db,
	}
}

func (r *pgsqlScheduleRepository) Create(ctx context.Context, schedule *domain.Schedule) (err error) {
	query := "INSERT INTO schedules (movie_id, theater_id, show_date, start_time, end_time) VALUES ($1, $2, $3, $4, $5)"
	_, err = r.db.ExecContext(
		ctx,
		query,
		schedule.MovieID,
		schedule.TheaterID,
		schedule.ShowDate,
		schedule.StartTime,
		schedule.EndTime,
	)
	return err
}

func (r *pgsqlScheduleRepository) Fetch(ctx context.Context) (schedules []domain.Schedule, err error) {
	query := "SELECT id, movie_id, theater_id, show_date, start_time, end_time, created_at, updated_at FROM schedules"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return schedules, err
	}

	defer rows.Close()

	for rows.Next() {
		var schedule domain.Schedule
		err := rows.Scan(&schedule.ID, &schedule.MovieID, &schedule.TheaterID, &schedule.ShowDate, &schedule.StartTime, &schedule.EndTime, &schedule.CreatedAt, &schedule.EndTime)

		if err != nil {
			return schedules, err
		}

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (r *pgsqlScheduleRepository) GetByID(ctx context.Context, id int64) (schedule domain.Schedule, err error) {
	query := "SELECT id, movie_id, theater_id, show_date, start_time, end_time, created_at, updated_at FROM schedules WHERE id = $1"
	err = r.db.QueryRowContext(ctx, query, id).Scan(
		&schedule.ID,
		&schedule.MovieID,
		&schedule.TheaterID,
		&schedule.ShowDate,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.CreatedAt,
		&schedule.UpdatedAt,
	)

	return
}

func (r *pgsqlScheduleRepository) Update(ctx context.Context, schedule *domain.Schedule) (err error) {
	query := "UPDATE schedules SET show_date = $1, start_time = $2, end_time = $3, updated_at = $4 WHERE id = $5"
	res, err := r.db.ExecContext(ctx, query, schedule.ShowDate, schedule.StartTime, schedule.EndTime, schedule.UpdatedAt, schedule.ID)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("something went wrong, rows affected: %d", affect)
	}

	return
}

func (r *pgsqlScheduleRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM schedules WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("something went wrong, rows affected: %d", affect)
	}

	return
}
