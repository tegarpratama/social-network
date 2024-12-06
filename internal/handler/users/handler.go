package users

import (
	"context"
	"social-network/internal/model/users"

	"github.com/gin-gonic/gin"
)

type userService interface {
	Register(ctx context.Context, req users.UserRegister) (int, error)
}

type handler struct {
	api     *gin.Engine
	userSvc userService
}

func NewHandler(api *gin.Engine, userSvc userService) *handler {
	return &handler{
		api:     api,
		userSvc: userSvc,
	}
}

func (h *handler) RouteList() {
	route := h.api.Group("/auth")
	route.POST("/register", h.RegisterUser)
}
