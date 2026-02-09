// internal/auth/infrastructure/repository/postgres.go
package repository

import (
	"context"
	"errors"
	"fmt"

	"api/auth/internal/auth/domain"
	"api/auth/internal/auth/infrastructure/repository/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	q *db.Queries
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{q: db.New(pool)}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := r.q.CreateUser(ctx, db.CreateUserParams{
		ID:           user.ID.String(),
		Email:        user.Email.String(),
		PasswordHash: user.PasswordHash,
		Role:         string(user.Role),
		Verified:     user.Verified,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	})
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	row, err := r.q.GetUserByEmail(ctx, email.String())
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return toDomain(row), nil
}

func (r *PostgresUserRepository) ExistsByEmail(ctx context.Context, email domain.Email) (bool, error) {
	exists, err := r.q.ExistsUserByEmail(ctx, email.String())
	if err != nil {
		return false, fmt.Errorf("failed to check email: %w", err)
	}
	return exists, nil
}

// toDomain maps sqlc-generated struct to domain entity
func toDomain(row db.User) *domain.User {
	return domain.ReconstructUser(
		domain.UserID(row.ID),
		domain.Email(row.Email),
		row.PasswordHash,
		domain.Role(row.Role),
		row.Verified,
		row.CreatedAt,
		row.UpdatedAt,
	)
}
