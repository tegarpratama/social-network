package posts

import (
	"net/http"
	"social-network/internal/model/posts"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListPosts(c *gin.Context) {
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

	result, err := h.postService.ListPosts(ctx, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := result.Data
	if len(data) == 0 {
		data = []posts.PostObj{}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "list posts",
		"paginate": result.Paginate,
		"data":     data,
	})
}
