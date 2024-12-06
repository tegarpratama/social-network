package categories

import (
	"context"
	"errors"
	"log"
	"net/http"
	"social-network/internal/helper"
	"social-network/internal/model/categories"
	"time"
)

func (s *service) UpdateCategory(ctx context.Context, categoryID int64, req *categories.CreateUpdateCategoryReq) *helper.Error {
	data, err := s.categoryRepo.CategoryDetail(ctx, categoryID)
	log.Println(data)
	if err != nil {
		return &helper.Error{
			Code:    http.StatusNotFound,
			Message: errors.New("category not found"),
		}
	}

	now := time.Now()
	model := categories.Category{
		Name:      req.Name,
		UpdatedAt: now,
	}

	err = s.categoryRepo.UpdateCategory(ctx, categoryID, model)
	if err != nil {
		return &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return nil
}
