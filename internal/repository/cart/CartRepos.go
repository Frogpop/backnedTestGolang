package cart

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CartRepos interface {
	CheckItem(cart_id, product_id uint64) (*models.CartItem, error)
	CreateItem(*models.CartItem) error
	UpdateItem(*models.CartItem) error
	RemoveItem(cartID, productID uint64) error
	GetCart(userID uint64) (*models.Cart, error)
	GetCartItems(userID uint64) (*[]dto.ItemDTO, error)
	ClearCart(userID uint64) error
	//QuantityCheckItem(cartID, productID uint64) (int, error)
}

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) CartRepos {
	return &cartRepo{db: db}
}

func (r *cartRepo) QuantityCheckItem(cartID, productID uint64) (int, error) {
	const op = "CartRepo.QuantityCheckItem"

	var cartItem *models.CartItem
	if err := r.db.Where("cart_id = ? and product_id = ?", cartID, productID).First(&cartItem).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return cartItem.Quantity, nil
}

func (r *cartRepo) CheckItem(cart_id, product_id uint64) (*models.CartItem, error) {
	const op = "CartRepo.CheckItem"

	var cartItem *models.CartItem
	if err := r.db.Where("cart_id = ? and product_id = ?", cart_id, product_id).First(&cartItem).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return cartItem, nil
}

func (r *cartRepo) CreateItem(item *models.CartItem) error {
	const op = "CartRepo.CreateItem"

	if err := r.db.Create(item).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *cartRepo) UpdateItem(item *models.CartItem) error {
	const op = "CartRepo.UpdateItem"

	if err := r.db.Where("cart_id = ? and product_id = ?", item.CartID, item.ProductID).Updates(&item).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *cartRepo) RemoveItem(cartID, productID uint64) error {
	const op = "CartRepo.RemoveItem"

	if err := r.db.Where("cart_id = ? and product_id = ?", cartID, productID).Delete(&models.CartItem{}).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *cartRepo) GetCart(userID uint64) (*models.Cart, error) {
	const op = "CartRepo.GetCart"
	var cart *models.Cart
	if err := r.db.Where("user_id = ?", userID).Preload("Items").First(&cart).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return cart, nil
}

func (r *cartRepo) GetCartItems(userID uint64) (*[]dto.ItemDTO, error) {
	const op = "CartRepo.GetCartItems"

	var items *[]dto.ItemDTO
	if err := r.db.Table("cart_items").Select("name", "price", "quantity").Joins("join products on products.id = cart_items.product_id").Where("cart_id = ?", userID).Scan(&items).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
}

func (r *cartRepo) ClearCart(cartID uint64) error {
	const op = "CartRepo.ClearCart"

	if err := r.db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
