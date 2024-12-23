package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()
	postIDStr := c.Param("postID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	errCode, err := h.postService.DeletePost(ctx, int64(postID), userID)
	if err != nil {
		c.JSON(errCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully deleted post",
	})
}
