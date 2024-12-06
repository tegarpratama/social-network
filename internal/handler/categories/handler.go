package categories

import (
	"context"
	"social-network/internal/model/categories"

	"github.com/gin-gonic/gin"
)

type categoryService interface {
	CreateCategory(ctx context.Context, req *categories.CreateUpdateCategoryReq) error
	DeleteCategory(ctx context.Context, categoryID int64) (int, error)
	UpdateCategory(ctx context.Context, categoryID int64, req *categories.CreateUpdateCategoryReq) (int, error)
	DetailCategory(ctx context.Context, categoryID int64) (*categories.CategoryObj, int, error)
	ListCategory(ctx context.Context, pageSize, pageIndex int) (*categories.ListCategoriesRes, error)
}

type handler struct {
	api         *gin.Engine
	categorySvc categoryService
}

func NewHandler(api *gin.Engine, categorySvc categoryService) *handler {
	return &handler{
		api:         api,
		categorySvc: categorySvc,
	}
}

func (h *handler) RouteList() {
	route := h.api.Group("/categories")
	route.GET("/", h.ListCategory)
	route.POST("/", h.CreateCategory)
	route.DELETE("/:categoryID", h.DeleteCtegory)
	route.PUT("/:categoryID", h.UpdateCategory)
	route.GET("/:categoryID", h.DetailCategory)
}
