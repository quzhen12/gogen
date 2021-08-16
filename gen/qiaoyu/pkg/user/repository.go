package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, email, password string) error
	FindUserByUsername(ctx context.Context, username string) (*User, error)
	DeleteUserById(ctx context.Context, id int) (int, error)
	UpdateUser(ctx context.Context, user *User) (int64, error)
}
