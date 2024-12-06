package categories

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/model/categories"
)

func (s *service) DetailCategory(ctx context.Context, categoryID int64) (*categories.CategoryObj, int, error) {
	category, err := s.categoryRepo.CategoryDetail(ctx, categoryID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if category == nil {
		return nil, http.StatusNotFound, errors.New("category not found")
	}

	return category, 0, nil
}
