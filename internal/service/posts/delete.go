package posts

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) DeletePost(ctx context.Context, postID int64, userID int64) (int, error) {
	post, err := s.postRepo.PostDetail(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	// Validation
	if post == nil {
		return http.StatusNotFound, errors.New("post not found")
	}

	if post.User.UserID != userID {
		return http.StatusBadRequest, errors.New("cannot delete this post")
	}

	now := time.Now()
	err = s.postRepo.SoftDeletePost(ctx, postID, now)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return 0, err
}
