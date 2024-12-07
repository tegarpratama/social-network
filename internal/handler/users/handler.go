package users

import (
	"context"
	"social-network/internal/middleware"
	"social-network/internal/model/users"

	"github.com/gin-gonic/gin"
)

type userService interface {
	Register(ctx context.Context, req users.UserRegisterReq) (int, error)
	Login(ctx context.Context, req users.UserLoginReq) (*users.UserLoginRes, int, error)
	RefreshToken(ctx context.Context, userID int64, req users.RefreshTokenReq) (*users.RefreshTokenRes, error)
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
	route.POST("/login", h.Login)

	routeRefresh := h.api.Group("/auth")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.RefreshToken)
}
