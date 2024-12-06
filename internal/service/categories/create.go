package categories

import (
	"context"
	"net/http"
	"social-network/internal/helper"
	"social-network/internal/model/categories"
	"time"
)

func (s *service) CreateCategory(ctx context.Context, req *categories.CreateUpdateCategoryReq) *helper.Error {
	now := time.Now()
	model := categories.Category{
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.categoryRepo.InsertCategory(ctx, &model)
	if err != nil {
		return &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return nil
}
