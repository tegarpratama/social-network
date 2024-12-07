package users

import (
	"net/http"
	"social-network/internal/model/users"

	"github.com/gin-gonic/gin"
)

func (h *handler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var req users.RefreshTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	result, err := h.userSvc.RefreshToken(ctx, userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messsage": "success generate new token",
		"data":     result,
	})
}
