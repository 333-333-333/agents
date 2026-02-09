// internal/auth/domain/port.go
package domain

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id UserID) (*User, error)
	FindByEmail(ctx context.Context, email Email) (*User, error)
	ExistsByEmail(ctx context.Context, email Email) (bool, error)
}
