package dto

type OrderDTO struct {
	OrderID uint64    `json:"order_id"`
	Status  string    `json:"status"`
	Items   []ItemDTO `json:"items"`
}

type UserOrdersDTO struct {
	UserID uint64     `json:"user_id" example:"1"`
	Orders []OrderDTO `json:"orders"`
}

type GetOrdersRequest struct {
	UserID uint64 `form:"user_id" binding:"required" example:"1"`
}

type ChangeOrderStatusRequest struct {
	OrderID uint64 `json:"order_id" binding:"required" example:"1"`
	Status  string `json:"status" binding:"required" example:"created"`
}
