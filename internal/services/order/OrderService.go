package order

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/mapper"
	"backnedTestGolang/internal/repository/order"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrders(userID uint64) (*dto.UserOrdersDTO, error)
	ChangeOrderStatus(orderID uint64, status string) error
}

type OrderServiceImpl struct {
	orderRepo order.OrderRepos
}

func NewOrderService(orderRepo order.OrderRepos) OrderService {
	return &OrderServiceImpl{orderRepo: orderRepo}
}

func (s *OrderServiceImpl) GetOrders(userID uint64) (*dto.UserOrdersDTO, error) {
	const op = "OrderServiceImpl.GetOrders"

	var orders *[]order.OrderWithItemsRaw
	var err error

	if orders, err = s.orderRepo.GetOrders(userID); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return mapper.ToUserOrdersDTO(orders), nil
}

func (s *OrderServiceImpl) ChangeOrderStatus(orderID uint64, status string) error {
	const op = "OrderServiceImpl.ChangeOrderStatus"

	if err := s.orderRepo.ChangeOrderStatus(orderID, status); errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
