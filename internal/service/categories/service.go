package categories

import (
	"context"
	"social-network/internal/model/categories"
	"time"
)

type categoryRepository interface {
	InsertCategory(ctx context.Context, model *categories.Category) error
	SoftDeleteCategory(ctx context.Context, categoryID int64, now time.Time) error
	CategoryDetail(ctx context.Context, categoryID int64) (*categories.CategoryObj, error)
	UpdateCategory(ctx context.Context, categoryID int64, model categories.Category) error
	ListCategory(ctx context.Context, limit, offset int) (*categories.ListCategoriesRes, error)
	TotalCategory(ctx context.Context) (int, error)
}

type service struct {
	categoryRepo categoryRepository
}

func NewService(categoryRepo categoryRepository) *service {
	return &service{
		categoryRepo: categoryRepo,
	}
}
