package categories

import (
	"net/http"
	"social-network/internal/helper"
	"social-network/internal/model/categories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateCategory(c *gin.Context) {
	ctx := c.Request.Context()

	categoryIDStr := c.Param("categoryID")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		helper.SetResponse(c, http.StatusInternalServerError, &helper.Response{
			Message: err.Error(),
		})
		return
	}

	var request categories.CreateUpdateCategoryReq
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.SetResponse(c, http.StatusBadRequest, &helper.Response{
			Message: err.Error(),
		})
		return
	}

	errService := h.categorySvc.UpdateCategory(ctx, categoryID, &request)
	if errService != nil {
		helper.SetResponse(c, errService.Code, &helper.Response{
			Message: errService.Message.Error(),
		})
		return
	}

	helper.SetResponse(c, http.StatusOK, &helper.Response{
		Message: "successfully updated category",
	})
}
