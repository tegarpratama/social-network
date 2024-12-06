package categories

import (
	"net/http"
	"social-network/internal/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteCtegory(c *gin.Context) {
	ctx := c.Request.Context()

	categoryIDStr := c.Param("categoryID")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		helper.SetResponse(c, http.StatusInternalServerError, &helper.Response{
			Message: err.Error(),
		})
		return
	}

	if err := h.categorySvc.DeleteCategory(ctx, categoryID); err != nil {
		helper.SetResponse(c, http.StatusInternalServerError, &helper.Response{
			Message: err.Message.Error(),
		})
		return
	}

	helper.SetResponse(c, http.StatusOK, &helper.Response{
		Message: "successfully deleted category",
	})
}
