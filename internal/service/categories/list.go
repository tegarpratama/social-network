package categories

import (
	"context"
	"math"
	"social-network/internal/model/categories"
	"social-network/internal/model/paginate"

	"github.com/rs/zerolog/log"
)

func (s *service) ListCategory(ctx context.Context, limit, page int) (*categories.ListCategoriesRes, error) {
	offset := (page - 1) * limit
	result, err := s.categoryRepo.ListCategory(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("")
		return result, nil

	}
	total, err := s.categoryRepo.TotalCategory(ctx)
	if err != nil {
		log.Error().Err(err).Msg("")
		return result, err
	}

	totalPages := math.Ceil(float64(total) / float64(limit))

	result.Paginate = paginate.Paginate{
		Limit:       limit,
		TotalPage:   int(totalPages),
		CurrentPage: page,
	}

	return result, nil
}
