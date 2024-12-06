package categories

import (
	"net/http"
	"social-network/internal/helper"
	"social-network/internal/model/categories"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCategory(c *gin.Context) {
	ctx := c.Request.Context()

	var request categories.CreateUpdateCategoryReq
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.SetResponse(c, http.StatusBadRequest, &helper.Response{
			Message: err.Error(),
		})
		return
	}

	if err := h.categorySvc.CreateCategory(ctx, &request); err != nil {
		helper.SetResponse(c, err.Code, &helper.Response{
			Message: err.Message.Error(),
		})
		return
	}

	helper.SetResponse(c, http.StatusCreated, &helper.Response{
		Message: "successfully created new category",
	})
}
