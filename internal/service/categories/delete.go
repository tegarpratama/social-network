package categories

import (
	"context"
	"net/http"
	"social-network/internal/helper"
	"time"
)

func (s *service) DeleteCategory(ctx context.Context, categoryID int64) *helper.Error {
	now := time.Now()
	err := s.categoryRepo.SoftDeleteCategory(ctx, categoryID, now)
	if err != nil {
		return &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return nil
}
