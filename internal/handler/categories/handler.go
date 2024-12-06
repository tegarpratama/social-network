package categories

import (
	"context"
	"social-network/internal/helper"
	"social-network/internal/model/categories"

	"github.com/gin-gonic/gin"
)

type categoryService interface {
	CreateCategory(ctx context.Context, req *categories.CreateUpdateCategoryReq) *helper.Error
	DeleteCategory(ctx context.Context, categoryID int64) *helper.Error
	UpdateCategory(ctx context.Context, categoryID int64, req *categories.CreateUpdateCategoryReq) *helper.Error
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
	route.POST("/", h.CreateCategory)
	route.DELETE("/:categoryID", h.DeleteCtegory)
	route.PUT("/:categoryID", h.UpdateCategory)
}
