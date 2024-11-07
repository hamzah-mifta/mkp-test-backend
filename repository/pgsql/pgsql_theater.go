package pgsql

import (
	"context"
	"database/sql"
	"fmt"
)

type pgsqlTheaterRepository struct {
	db *sql.DB
}

func NewPgsqlTheaterRepository(db *sql.DB) *pgsqlTheaterRepository {
	return &pgsqlTheaterRepository{
		db: db,
	}
}

func (r *pgsqlTheaterRepository) IsExist(ctx context.Context, key string, val interface{}) (isExist bool, err error) {
	allowedKeys := map[string]bool{
		"id":      true,
		"name":    true,
		"city":    true,
		"address": true,
	}

	if !allowedKeys[key] {
		return false, fmt.Errorf("invalid key: %s", key)
	}

	var exist int

	query := fmt.Sprintf("SELECT 1 FROM theaters WHERE %s = $1", key)

	err = r.db.QueryRowContext(ctx, query, val).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows found, meaning the movie doesn't exist.
			return false, nil
		}

		return false, err
	}

	return true, nil
}
