package users

import (
	"context"
	"social-network/internal/model/users"
)

type userRepository interface {
	UsernameExist(ctx context.Context, username string) (bool, error)
	InsertUser(ctx context.Context, model users.User) error
}

type service struct {
	userRepo userRepository
}

func NewService(userRepo userRepository) *service {
	return &service{
		userRepo: userRepo,
	}
}
