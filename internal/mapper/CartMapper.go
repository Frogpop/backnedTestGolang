package mapper

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
)

/*func ToCartDTO(cart *models.Cart) *dto.CartDTO {
	items := make([]*dto.CartItemDTO, len(cart.Items))
	for i, item := range cart.Items {
		items[i] = ToCartItemDTO(&item)
	}
	return &dto.CartDTO{
		UserID: cart.UserID,
		Items:  items,
	}
}*/

func ToCartItemDTO(item *models.CartItem) *dto.CartItemDTO {
	return &dto.CartItemDTO{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
}
