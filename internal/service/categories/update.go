package categories

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/model/categories"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) UpdateCategory(ctx context.Context, categoryID int64, req *categories.CreateUpdateCategoryReq) (int, error) {
	category, err := s.categoryRepo.CategoryDetail(ctx, categoryID)
	if err != nil {
		log.Error().Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	if category == nil {
		return http.StatusNotFound, errors.New("category not found")
	}

	now := time.Now()
	model := categories.Category{
		Name:      req.Name,
		UpdatedAt: now,
	}

	err = s.categoryRepo.UpdateCategory(ctx, categoryID, model)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}
