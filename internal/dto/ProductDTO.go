package dto

type ProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ItemDTO struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
