package dto

type ProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ItemDTO struct {
	Name     string  `json:"name" example:"ProductName"`
	Price    float64 `json:"price" example:"100.0"`
	Quantity int     `json:"quantity" example:"1"`
}
