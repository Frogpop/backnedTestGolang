package repository

import (
	"backnedTestGolang/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type OrderWithItemsRaw struct {
	UserID   uint64
	OrderID  uint64
	Status   string
	Name     string
	Price    float64
	Quantity int
}

type OrderRepos interface {
	CreateOrder(order *models.Order) (uint64, error)
	GetOrders(userID uint64) (*[]OrderWithItemsRaw, error)
	ChangeOrderStatus(orderID uint64, status string) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepos {
	return &orderRepo{db}
}

func (r *orderRepo) CreateOrder(order *models.Order) (uint64, error) {
	const op = "OrderRepo.CreateOrder"
	if err := r.db.Create(&order).Error; err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return order.ID, nil
}

func (r *orderRepo) GetOrders(userID uint64) (*[]OrderWithItemsRaw, error) {
	const op = "OrderRepo.GetOrders"

	var userOrders []OrderWithItemsRaw
	if err := r.db.Table("orders").Select("user_id, order_id, status, name, price, quantity").Joins("join order_items oi on orders.id = oi.order_id join products p on oi.product_id = p.id").Where("user_id = ?", userID).Scan(&userOrders).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &userOrders, nil
}

func (r *orderRepo) ChangeOrderStatus(orderID uint64, status string) error {
	const op = "OrderRepo.ChangeOrderStatus"

	var order models.Order
	if err := r.db.Where("id = ?", orderID).First(&order).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	order.Status = status
	return r.db.Save(&order).Error
}
