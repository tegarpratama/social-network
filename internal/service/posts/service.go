package posts

import (
	"context"
	"social-network/internal/model/posts"
	"time"
)

type postRepository interface {
	InsertPost(ctx context.Context, model *posts.Post) error
	UpdatePost(ctx context.Context, postID int64, model *posts.Post) error
	PostDetail(ctx context.Context, postID int64) (*posts.PostObj, error)
	SoftDeletePost(ctx context.Context, postID int64, now time.Time) error
	ListPosts(ctx context.Context, limit, offset int) (*[]posts.PostObj, error)
	TotalPosts(ctx context.Context) (int, error)
}

type service struct {
	postRepo postRepository
}

func NewService(postRepo postRepository) *service {
	return &service{
		postRepo: postRepo,
	}
}
