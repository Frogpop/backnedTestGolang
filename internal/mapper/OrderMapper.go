package mapper

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/repository/order"
)

func ToUserOrdersDTO(orders_raw *[]order.OrderWithItemsRaw) *dto.UserOrdersDTO {
	userOrders := &dto.UserOrdersDTO{UserID: (*orders_raw)[0].UserID}
	ordersMap := make(map[uint64]dto.OrderDTO)
	for _, row := range *orders_raw {
		order, exists := ordersMap[row.OrderID]
		if !exists {
			ordersMap[row.OrderID] = dto.OrderDTO{
				OrderID: row.OrderID,
				Status:  row.Status,
				Items:   []dto.ItemDTO{},
			}
			order, _ = ordersMap[row.OrderID]
		}
		order.Items = append(order.Items, dto.ItemDTO{
			Name:     row.Name,
			Price:    row.Price,
			Quantity: row.Quantity,
		})

		ordersMap[row.OrderID] = order
	}

	for _, o := range ordersMap {
		userOrders.Orders = append(userOrders.Orders, o)
	}

	return userOrders
}
