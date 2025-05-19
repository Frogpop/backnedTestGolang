package cart

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
	"backnedTestGolang/internal/repository/cart"
	"backnedTestGolang/internal/repository/order"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CartService interface {
	AddProduct(cartID, productID uint64, quantity int) error
	ReduceProduct(cartID, productID uint64, quantity int) error
	RemoveProduct(cartID, ProductID uint64) error
	GetCartItems(userID uint64) (*[]dto.ItemDTO, error)
	Checkout(userID uint64) error
}

type cartServiceImpl struct {
	cartRepo  cart.CartRepos
	orderRepo order.OrderRepos
}

func NewCartService(cartRepo cart.CartRepos, orderRepo order.OrderRepos) CartService {
	return &cartServiceImpl{cartRepo: cartRepo, orderRepo: orderRepo}
}

func (s *cartServiceImpl) AddProduct(cartID, productID uint64, quantity int) error {
	const op = "cartServiceImpl.AddProduct"
	var item *models.CartItem
	var err error
	if item, err = s.cartRepo.CheckItem(cartID, productID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s: %w", op, err)
	}

	if item == nil {
		if err = s.cartRepo.CreateItem(&models.CartItem{CartID: cartID, ProductID: productID, Quantity: quantity}); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	}
	item.Quantity += quantity

	if err = s.cartRepo.UpdateItem(item); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *cartServiceImpl) ReduceProduct(cartID, productID uint64, quantity int) error {
	const op = "cartServiceImpl.ReduceProduct"

	var item *models.CartItem
	var err error
	if item, err = s.cartRepo.CheckItem(cartID, productID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s: %w", op, err)
	}
	if item == nil {
		return nil
	} else if item.Quantity-quantity <= 0 {
		if err = s.cartRepo.RemoveItem(cartID, productID); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	}
	item.Quantity -= quantity
	if err = s.cartRepo.UpdateItem(item); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *cartServiceImpl) RemoveProduct(cartID, ProductID uint64) error {
	const op = "cartServiceImpl.RemoveProduct"

	if err := s.cartRepo.RemoveItem(cartID, ProductID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *cartServiceImpl) GetCartItems(userID uint64) (*[]dto.ItemDTO, error) {
	const op = "cartServiceImpl.GetCartItems"

	var items *[]dto.ItemDTO
	var err error
	if items, err = s.cartRepo.GetCartItems(userID); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
}

func (s *cartServiceImpl) Checkout(userID uint64) error {
	const op = "cartServiceImpl.Checkout"

	cart, err := s.cartRepo.GetCart(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if cart.Items == nil {
		return nil
	}

	if err = s.cartRepo.ClearCart(userID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	orderItems := make([]models.OrderItem, len(cart.Items))
	for i, item := range cart.Items {
		orderItems[i] = models.OrderItem{ProductID: item.ProductID, Quantity: item.Quantity}
	}

	_, err = s.orderRepo.CreateOrder(&models.Order{UserID: cart.UserID, Items: orderItems})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
