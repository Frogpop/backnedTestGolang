package mocks

import (
	"backnedTestGolang/internal/dto"
	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) GetUserOrders(userID uint64) (*dto.UserOrdersDTO, error) {
	args := m.Called(userID)
	return args.Get(0).(*dto.UserOrdersDTO), args.Error(1)
}

func (m *MockOrderService) ChangeOrderStatus(orderID uint64, status string) error {
	args := m.Called(orderID, status)
	return args.Error(0)
}
