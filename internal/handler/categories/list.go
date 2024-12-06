package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListCategory(c *gin.Context) {
	ctx := c.Request.Context()
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limitStr == "" {
		limit = 5
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || pageStr == "" {
		page = 1
	}

	result, err := h.categorySvc.ListCategory(ctx, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "list categories",
		"paginate": result.Paginate,
		"data":     result.Data,
	})
}
