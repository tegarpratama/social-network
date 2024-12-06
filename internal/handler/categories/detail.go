package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DetailCategory(c *gin.Context) {
	ctx := c.Request.Context()

	categoryIDStr := c.Param("categoryID")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	category, errCode, err := h.categorySvc.DetailCategory(ctx, int64(categoryID))
	if err != nil {
		c.JSON(errCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "detail category",
		"data":    category,
	})
}
