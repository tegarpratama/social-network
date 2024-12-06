package users

import (
	"net/http"
	"social-network/internal/model/users"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req users.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	errCode, err := h.userSvc.Register(ctx, req)
	if err != nil {
		c.JSON(errCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register successfully",
	})
}
