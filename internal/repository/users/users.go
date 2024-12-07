package users

import (
	"context"
	"database/sql"
	"social-network/internal/model/users"
	"time"
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

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*users.User, error) {
	query := `SELECT id, username, password, created_at, updated_at FROM users WHERE username = ? AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, username)

	var user users.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*users.RefreshToken, error) {
	query := `
		SELECT id, user_id, refresh_token, expired_at, created_at 
		FROM refresh_tokens 
		WHERE user_id = ? AND expired_at >= ?`

	var refreshToken users.RefreshToken
	row := r.db.QueryRowContext(ctx, query, userID, now)
	err := row.Scan(&refreshToken.ID, &refreshToken.UserID, &refreshToken.RefreshToken, &refreshToken.ExpiredAt, &refreshToken.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &refreshToken, nil
}

func (r *repository) InsertRefreshToken(ctx context.Context, model *users.RefreshToken) error {
	query := `
		INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at)
		VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt)

	return err
}

func (r *repository) GetUserByID(ctx context.Context, userID int64) (*users.User, error) {
	query := `SELECT id, username, password, created_at, updated_at FROM users WHERE id = ? AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, userID)

	var user users.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) DeleteRefreshToken(ctx context.Context, userID int64) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = ?`
	_, err := r.db.ExecContext(ctx, query, userID)

	return err
}
