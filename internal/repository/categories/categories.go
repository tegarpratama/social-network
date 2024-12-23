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

	return err
}

func (r *repository) ListCategory(ctx context.Context, limit, offset int) (*[]categories.CategoryObj, error) {

	query := `SELECT id, name, created_at, updated_at
		FROM categories 
		WHERE deleted_at IS NULL
		ORDER BY id DESC
		LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		data []categories.CategoryObj
	)

	for rows.Next() {
		var categoryTemp categories.CategoryObj
		err := rows.Scan(&categoryTemp.ID, &categoryTemp.Name, &categoryTemp.CreatedAt, &categoryTemp.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, categories.CategoryObj{
			ID:        categoryTemp.ID,
			Name:      categoryTemp.Name,
			CreatedAt: categoryTemp.CreatedAt,
			UpdatedAt: categoryTemp.UpdatedAt,
		})
	}

	return &data, nil
}

func (r *repository) TotalCategory(ctx context.Context) (int, error) {
	total := 0

	query := `SELECT COUNT(id) as total FROM categories WHERE deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query)
	err := row.Scan(&total)

	return total, err
}
