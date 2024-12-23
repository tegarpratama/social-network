package posts

import (
	"context"
	"math"
	"social-network/internal/model/paginate"
	"social-network/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) ListPosts(ctx context.Context, limit, page int) (*posts.ListPostsRes, error) {
	offset := (page - 1) * limit
	listPosts, err := s.postRepo.ListPosts(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, nil

	}

	total, err := s.postRepo.TotalPosts(ctx)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, nil
	}

	totalPages := math.Ceil(float64(total) / float64(limit))

	result := posts.ListPostsRes{
		Paginate: paginate.Paginate{
			Limit:       limit,
			TotalPage:   int(totalPages),
			CurrentPage: page,
		},
		Data: *listPosts,
	}

	return &result, nil
}
