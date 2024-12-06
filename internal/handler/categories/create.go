package categories

import (
	"net/http"
	"social-network/internal/model/categories"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCategory(c *gin.Context) {
	ctx := c.Request.Context()

	var request categories.CreateUpdateCategoryReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.categorySvc.CreateCategory(ctx, &request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created new category",
	})
}
