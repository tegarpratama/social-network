package users

import (
	"net/http"
	"social-network/internal/model/users"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req users.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, errCOde, err := h.userSvc.Login(ctx, req)
	if err != nil {
		c.JSON(errCOde, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully login",
		"data":    result,
	})
}
