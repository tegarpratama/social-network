package posts

import (
	"context"
	"database/sql"
	"errors"
	"social-network/internal/model/posts"
	"time"
)

type queryPostTemp struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       int64     `json:"user_id"`
	Username     string    `json:"username"`
	CategoryID   int64     `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

func (r *repository) InsertPost(ctx context.Context, model *posts.Post) error {
	query := `
    INSERT INTO posts (user_id, category_id, title, content, created_at, updated_at)
    VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.CategoryID, model.Title, model.Content, model.CreatedAt, model.UpdatedAt)

	return err
}

func (r *repository) UpdatePost(ctx context.Context, postID int64, model *posts.Post) error {
	query := `UPDATE posts SET category_id = ?, title = ?, content = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, model.CategoryID, model.Title, model.Content, model.UpdatedAt, postID)
	return err
}

func (r *repository) PostDetail(ctx context.Context, postID int64) (*posts.PostObj, error) {
	query := `
		SELECT 
			p.id, 
			p.title, 
			p.content, 
			p.created_at, 
			p.updated_at,
			u.id as user_id,
			u.username,
			c.id as category_id,
			c.name as category_name
		FROM posts as p
		JOIN categories as c ON c.id = p.category_id
		JOIN users as u ON u.id = p.user_id
		WHERE p.id = ?
		AND p.deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, postID)

	var (
		postObj  posts.PostObj
		postTemp queryPostTemp
	)

	err := row.Scan(&postTemp.ID, &postTemp.Title, &postTemp.Content, &postTemp.CreatedAt, &postTemp.UpdatedAt, &postTemp.UserID, &postTemp.Username, &postTemp.CategoryID, &postTemp.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	postObj = posts.PostObj{
		ID:      postTemp.ID,
		Title:   postTemp.Title,
		Content: postTemp.Content,
		User: posts.UserObj{
			UserID:   postTemp.UserID,
			Username: postTemp.Username,
		},
		Category: posts.CategoryObj{
			CategoryID: postTemp.CategoryID,
			Name:       postTemp.CategoryName,
		},
		CreatedAt: postTemp.CreatedAt,
		UpdatedAt: postTemp.UpdatedAt,
	}

	return &postObj, nil
}

func (r *repository) SoftDeletePost(ctx context.Context, postID int64, now time.Time) error {
	query := `UPDATE posts SET deleted_at = ? WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, now, postID)
	if err != nil {
		return err
	}

	countAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func (r *repository) ListPosts(ctx context.Context, limit, offset int) (*[]posts.PostObj, error) {
	query := `
		SELECT 
			p.id, 
			p.title, 
			p.content, 
			p.created_at, 
			p.updated_at,
			u.id as user_id,
			u.username,
			c.id as category_id,
			c.name as category_name
		FROM posts as p
		JOIN categories as c ON c.id = p.category_id
		JOIN users as u ON u.id = p.user_id
		WHERE p.deleted_at IS NULL
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		listPosts []posts.PostObj
	)

	for rows.Next() {
		var postTemp queryPostTemp
		err := rows.Scan(&postTemp.ID, &postTemp.Title, &postTemp.Content, &postTemp.CreatedAt, &postTemp.UpdatedAt, &postTemp.UserID, &postTemp.Username, &postTemp.CategoryID, &postTemp.CategoryName)

		if err != nil {
			return nil, err
		}

		listPosts = append(listPosts, posts.PostObj{
			ID:      postTemp.ID,
			Title:   postTemp.Title,
			Content: postTemp.Content,
			User: posts.UserObj{
				UserID:   postTemp.UserID,
				Username: postTemp.Username,
			},
			Category: posts.CategoryObj{
				CategoryID: postTemp.CategoryID,
				Name:       postTemp.CategoryName,
			},
			CreatedAt: postTemp.CreatedAt,
			UpdatedAt: postTemp.UpdatedAt,
		})
	}

	return &listPosts, nil
}

func (r *repository) TotalPosts(ctx context.Context) (int, error) {
	total := 0

	query := `SELECT COUNT(id) as total FROM posts WHERE deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query)
	err := row.Scan(&total)

	return total, err
}
