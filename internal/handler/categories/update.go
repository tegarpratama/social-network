package categories

import (
	"net/http"
	"social-network/internal/model/categories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateCategory(c *gin.Context) {
	ctx := c.Request.Context()

	categoryIDStr := c.Param("categoryID")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req categories.CreateUpdateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	errCode, err := h.categorySvc.UpdateCategory(ctx, int64(categoryID), &req)
	if err != nil {
		c.JSON(errCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully updated category",
	})
}
