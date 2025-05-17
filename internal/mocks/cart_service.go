package mocks

import (
	"backnedTestGolang/internal/dto"
	"github.com/stretchr/testify/mock"
)

type MockCartService struct {
	mock.Mock
}

func (m *MockCartService) AddProduct(cartID, productID uint64, quantity int) error {
	args := m.Called(cartID, productID, quantity)
	return args.Error(0)
}

func (m *MockCartService) ReduceProduct(cartID, productID uint64, quantity int) error {
	args := m.Called(cartID, productID, quantity)
	return args.Error(0)
}

func (m *MockCartService) RemoveProduct(cartID, productID uint64) error {
	args := m.Called(cartID, productID)
	return args.Error(0)
}

func (m *MockCartService) GetCartItems(cartID uint64) (*[]dto.ItemDTO, error) {
	args := m.Called(cartID)
	return args.Get(0).(*[]dto.ItemDTO), args.Error(1)
}

func (m *MockCartService) Checkout(userID uint64) error {
	args := m.Called(userID)
	return args.Error(0)
}
