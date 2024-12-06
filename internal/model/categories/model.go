package categories

import "time"

type (
	Category struct {
		ID        int64     `db:"id"`
		Name      string    `db:"name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		DeletedAt time.Time `db:"deleted_at"`
	}
)

type (
	CreateUpdateCategoryReq struct {
		Name string `json:"name"`
	}
)

type (
	CategoryObj struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CategoryRes struct {
		Data []CategoryObj `json:"data"`
	}
)