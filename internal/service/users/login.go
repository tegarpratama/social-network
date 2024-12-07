package users

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/model/users"
	"social-network/pkg/jwt"
	"social-network/pkg/refresh_token"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req users.UserLoginReq) (*users.UserLoginRes, int, error) {
	// Check user exist
	user, err := s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Err(err).Msg("")
		return nil, http.StatusInternalServerError, err
	}

	if user == nil {
		log.Error().Err(err).Msg("")
		return nil, http.StatusNotFound, errors.New("wrong username or password")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, http.StatusNotFound, errors.New("wrong username or password")
	}

	// Generate token
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.SECRET_JWT)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, http.StatusInternalServerError, err
	}

	//  Prepare value for response
	result := &users.UserLoginRes{
		ID:       int(user.ID),
		Username: user.Username,
		Token:    token,
	}

	// Get existing refresh token
	existingRefreshToken, err := s.userRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, http.StatusInternalServerError, err
	}

	if existingRefreshToken != nil {
		result.RefreshToken = existingRefreshToken.RefreshToken
		return result, http.StatusOK, nil
	}

	// Generate refresh token
	refreshToken := refresh_token.GenerateRefreshToken()
	if refreshToken == "" {
		return result, http.StatusInternalServerError, errors.New("failed to generate refresh token")
	}

	result.RefreshToken = refreshToken

	err = s.userRepo.DeleteRefreshToken(ctx, user.ID)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, 0, err
	}

	now := time.Now()
	err = s.userRepo.InsertRefreshToken(ctx, &users.RefreshToken{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    now,
	})

	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, 0, err
	}

	return result, 0, nil
}
