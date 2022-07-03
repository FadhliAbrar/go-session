package repository

import (
	"context"
	"github.com/FadhliAbrar/golang-session/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	FindUserById(ctx context.Context, Id int) model.User
	FindUserByUsername(ctx context.Context, username string) (model.User, error)
}
