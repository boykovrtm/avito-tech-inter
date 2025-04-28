package user

import "context"

type Repository interface {
	Add(ctx context.Context, user *User, hash string) error
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindHashByUsername(ctx context.Context, username string) (*string, error)
}
