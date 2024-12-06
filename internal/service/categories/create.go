package categories

import (
	"context"
	"social-network/internal/model/categories"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateCategory(ctx context.Context, req *categories.CreateUpdateCategoryReq) error {
	now := time.Now()
	model := categories.Category{
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.categoryRepo.InsertCategory(ctx, &model)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return err
}
