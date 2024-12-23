package posts

import (
	"social-network/internal/model/paginate"
	"time"
)

type (
	Post struct {
		ID         int64      `db:"id"`
		UserID     int64      `db:"user_id"`
		CategoryID int64      `db:"category_id"`
		Title      string     `db:"title"`
		Content    string     `db:"content"`
		CreatedAt  time.Time  `db:"created_at"`
		UpdatedAt  time.Time  `db:"updated_at"`
		DeletedAt  *time.Time `db:"deleted_at"`
	}
)

type (
	CreateUpdatePostReq struct {
		CategoryID int64  `json:"category_id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
	}
)

type (
	UserObj struct {
		UserID   int64  `json:"id"`
		Username string `json:"username"`
	}

	CategoryObj struct {
		CategoryID int64  `json:"id"`
		Name       string `json:"name"`
	}

	PostObj struct {
		ID        int64       `json:"id"`
		User      UserObj     `json:"user"`
		Category  CategoryObj `json:"category"`
		Title     string      `json:"title"`
		Content   string      `json:"content"`
		CreatedAt time.Time   `json:"created_at"`
		UpdatedAt time.Time   `json:"updated_at"`
	}

	ListPostsRes struct {
		Paginate paginate.Paginate `json:"paginate"`
		Data     []PostObj         `json:"data"`
	}
)
