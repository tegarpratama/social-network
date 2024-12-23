package posts

import (
	"context"
	"social-network/internal/model/posts"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, req *posts.CreateUpdatePostReq, userID int64) error {
	now := time.Now()
	model := posts.Post{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Content:    req.Content,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err := s.postRepo.InsertPost(ctx, &model)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return err
}
