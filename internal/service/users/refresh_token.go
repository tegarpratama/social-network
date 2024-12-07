package users

import (
	"context"
	"errors"
	"social-network/internal/model/users"
	"social-network/pkg/jwt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) RefreshToken(ctx context.Context, userID int64, req users.RefreshTokenReq) (*users.RefreshTokenRes, error) {
	existingRefreshToken, err := s.userRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	if existingRefreshToken == nil {
		return nil, errors.New("refresh token has expired")
	}

	if existingRefreshToken.RefreshToken != req.RefreshToken {
		return nil, errors.New("refresh token is invalid")
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	token, err := jwt.CreateToken(userID, user.Username, s.cfg.SECRET_JWT)
	if err != nil {
		return nil, err
	}

	result := &users.RefreshTokenRes{
		Token: token,
	}

	return result, nil
}
