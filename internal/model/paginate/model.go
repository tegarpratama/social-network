package paginate

type (
	Paginate struct {
		Limit       int `json:"limit"`
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
	}
)
