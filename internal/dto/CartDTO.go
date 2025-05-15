package dto

type CartDTO struct {
	UserID uint64     `json:"user_id"`
	Items  *[]ItemDTO `json:"items"`
}

type CartItemDTO struct {
	ProductID uint64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AddProductRequest struct {
	CartID    uint64 `json:"cart_id" binding:"required"`
	ProductID uint64 `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type ReduceProductRequest struct {
	CartID    uint64 `json:"cart_id" binding:"required"`
	ProductID uint64 `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type RemoveProductRequest struct {
	CartID    uint64 `json:"cart_id" binding:"required"`
	ProductID uint64 `json:"product_id" binding:"required"`
}

type CheckoutRequest struct {
	UserID uint64 `json:"user_id" binding:"required"`
}

type GetCartRequest struct {
	UserID uint64 `form:"user_id" binding:"required"`
}
