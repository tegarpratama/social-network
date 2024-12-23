package posts

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/model/posts"
)

func (s *service) DetailPost(ctx context.Context, postID int64) (*posts.PostObj, int, error) {
	post, err := s.postRepo.PostDetail(ctx, postID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if post == nil {
		return nil, http.StatusNotFound, errors.New("category not found")
	}

	return post, 0, nil
}
