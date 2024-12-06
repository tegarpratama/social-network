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
)

type (
	UserRegister struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"password_confirm"`
	}
)
