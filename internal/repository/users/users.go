package users

import (
	"context"
	"database/sql"
	"social-network/internal/model/users"
)

func (r *repository) UsernameExist(ctx context.Context, username string) (bool, error) {
	var id int
	query := `SELECT id FROM users WHERE username = ? AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, username)
	err := row.Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *repository) InsertUser(ctx context.Context, model users.User) error {
	query := `
    INSERT INTO users (username, password, created_at, updated_at)
    VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Username, model.Password, model.CreatedAt, model.UpdatedAt)

	return err
}
