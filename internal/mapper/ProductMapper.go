package mapper

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
)

func ToProductDTO(p *models.Product) *dto.ProductDTO {
	return &dto.ProductDTO{
		Name:  p.Name,
		Price: p.Price,
	}
}
