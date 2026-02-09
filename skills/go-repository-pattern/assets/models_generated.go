// db/models.go (generated)
package db

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	Verified     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
