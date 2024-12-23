package posts

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/model/posts"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) UpdatePost(ctx context.Context, postID int64, req *posts.CreateUpdatePostReq, userID int64) (int, error) {
	now := time.Now()

	// Validation
	post, err := s.postRepo.PostDetail(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	if post == nil {
		return http.StatusNotFound, errors.New("post not found")
	}

	if post.User.UserID != userID {
		return http.StatusBadRequest, errors.New("cannot update this post")
	}

	model := posts.Post{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Content:    req.Content,
		UpdatedAt:  now,
	}

	err = s.postRepo.UpdatePost(ctx, postID, &model)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return 0, err
}
