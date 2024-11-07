package pgsql

import (
	"context"
	"database/sql"

	"github.com/hamzah-mifta/mkp-test-backend/domain"
)

type pgsqlUserRepository struct {
	db *sql.DB
}

func NewPgsqlUserRepository(db *sql.DB) *pgsqlUserRepository {
	return &pgsqlUserRepository{
		db: db,
	}
}

func (r *pgsqlUserRepository) Create(ctx context.Context, user domain.User) (err error) {
	query := "INSERT INTO users (email, password, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err = r.db.ExecContext(ctx, query, user.Email, user.Password, user.Name, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *pgsqlUserRepository) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	query := "SELECT id, email, password, name, created_at, updated_at FROM users WHERE email = $1"
	err = r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	return
}
