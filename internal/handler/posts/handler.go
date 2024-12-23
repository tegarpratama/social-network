package posts

import (
	"context"
	"social-network/internal/middleware"
	"social-network/internal/model/posts"

	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, req *posts.CreateUpdatePostReq, userID int64) error
	UpdatePost(ctx context.Context, postID int64, req *posts.CreateUpdatePostReq, userID int64) (int, error)
	DetailPost(ctx context.Context, postID int64) (*posts.PostObj, int, error)
	DeletePost(ctx context.Context, postID int64, userID int64) (int, error)
	ListPosts(ctx context.Context, limit, page int) (*posts.ListPostsRes, error)
}

type handler struct {
	api         *gin.Engine
	postService postService
}

func NewHandler(api *gin.Engine, postService postService) *handler {
	return &handler{
		api:         api,
		postService: postService,
	}
}

func (h *handler) RouteList() {
	route := h.api.Group("/posts")
	routeWithAuth := h.api.Group("/posts")

	routeWithAuth.Use(middleware.AuthMiddleware())

	routeWithAuth.POST("/", h.CreatePost)
	routeWithAuth.PUT("/:postID", h.UpdatePost)
	route.GET("/:postID", h.DetailPost)
	routeWithAuth.DELETE("/:postID", h.DeletePost)
	route.GET("/", h.ListPosts)
}
