package categories

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) DeleteCategory(ctx context.Context, categoryID int64) (int, error) {
	category, err := s.categoryRepo.CategoryDetail(ctx, categoryID)
	if err != nil {
		log.Error().Err(err).Msg("")
		return http.StatusInternalServerError, err
	}

	if category == nil {
		return http.StatusNotFound, errors.New("category not found")
	}

	now := time.Now()
	err = s.categoryRepo.SoftDeleteCategory(ctx, categoryID, now)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return 0, err
}
