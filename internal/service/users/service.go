package users

import (
	"context"
	"social-network/internal/config"
	"social-network/internal/model/users"
	"time"
)

type userRepository interface {
	UsernameExist(ctx context.Context, username string) (bool, error)
	InsertUser(ctx context.Context, model users.User) error
	GetUserByUsername(ctx context.Context, username string) (*users.User, error)
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*users.RefreshToken, error)
	InsertRefreshToken(ctx context.Context, model *users.RefreshToken) error
}

type service struct {
	userRepo userRepository
	cfg      *config.ConfigTypes
}

func NewService(userRepo userRepository, cgf *config.ConfigTypes) *service {
	return &service{
		userRepo: userRepo,
		cfg:      cgf,
	}
}
