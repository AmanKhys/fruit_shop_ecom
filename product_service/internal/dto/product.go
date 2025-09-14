package dto

type ProductUserResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductAdminResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	IsDelted bool    `json:"is_deleted"`
}
