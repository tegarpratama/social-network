package users

import "time"

type (
	User struct {
		ID        int64      `db:"id"`
		Username  string     `db:"username"`
		Password  string     `db:"password"`
		CreatedAt time.Time  `db:"created_at"`
		UpdatedAt time.Time  `db:"updated_at"`
		DeletedAt *time.Time `db:"deleted_at"`
	}

	RefreshToken struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		RefreshToken string    `db:"refresh_token"`
		ExpiredAt    time.Time `db:"expired_at"`
		CreatedAt    time.Time `db:"created_at"`
	}
)

type (
	UserRegisterReq struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"password_confirm"`
	}

	UserLoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type (
	UserLoginRes struct {
		ID           int    `json:"id"`
		Username     string `json:"username"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
