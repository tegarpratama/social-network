package posts

import (
	"net/http"
	"social-network/internal/model/posts"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()

	postIDStr := c.Param("postID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var request posts.CreateUpdatePostReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	errCode, err := h.postService.UpdatePost(ctx, int64(postID), &request, userID)
	if err != nil {
		c.JSON(errCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully updated posts",
	})
}
