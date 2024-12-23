package posts

import (
	"net/http"
	"social-network/internal/model/posts"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateUpdatePostReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	if err := h.postService.CreatePost(ctx, &request, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created posts",
	})
}
