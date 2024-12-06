package users

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/model/users"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, req users.UserRegister) (int, error) {
	exist, err := s.userRepo.UsernameExist(ctx, req.Username)
	if err != nil {
		log.Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	if exist {
		return http.StatusBadRequest, errors.New("username already exist")
	}

	if req.Password != req.PasswordConfirm {
		return http.StatusBadRequest, errors.New("password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	now := time.Now()
	model := users.User{
		Username:  req.Username,
		Password:  string(passwordHash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.userRepo.InsertUser(ctx, model)
	if err != nil {
		log.Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
