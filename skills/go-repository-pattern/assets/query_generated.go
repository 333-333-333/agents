// db/query.sql.go (generated)
package db

import "context"

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) { ... }
func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) { ... }
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) { ... }
func (q *Queries) ExistsUserByEmail(ctx context.Context, email string) (bool, error) { ... }
