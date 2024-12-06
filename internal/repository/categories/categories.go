package categories

import (
	"context"
	"database/sql"
	"errors"
	"social-network/internal/model/categories"
	"time"
)

func (r *repository) InsertCategory(ctx context.Context, model *categories.Category) error {
	query := `
		INSERT INTO categories (name, created_at, updated_at) 
		VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.Name, model.CreatedAt, model.UpdatedAt)
	return err
}

func (r *repository) SoftDeleteCategory(ctx context.Context, categoryID int64, now time.Time) error {
	query := `UPDATE categories SET deleted_at = ? WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, now, categoryID)
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

func (r *repository) CategoryDetail(ctx context.Context, categoryID int64) (*categories.CategoryObj, error) {
	query := `SELECT id, name, created_at, updated_at FROM categories WHERE id = ? AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, categoryID)

	var data categories.CategoryObj
	err := row.Scan(&data.ID, &data.Name, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &data, nil
}

func (r *repository) UpdateCategory(ctx context.Context, categoryID int64, model categories.Category) error {
	query := `UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, model.Name, model.UpdatedAt, categoryID)
	if err != nil {
		return err
	}

	return nil
}
