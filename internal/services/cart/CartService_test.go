package cart

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
	"backnedTestGolang/internal/repository/cart"
	"backnedTestGolang/internal/repository/order"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&models.Cart{}, &models.CartItem{}, &models.Product{}, &models.Order{}, &models.OrderItem{}); err != nil {
		panic("migration failed")
	}

	return db
}

func TestCartServiceImpl_AddProduct(t *testing.T) {
	db := setupTestDB(t)

	err := db.Create(&models.Product{Name: "Product A", Price: 100}).Error
	assert.NoError(t, err)

	cartRepo := cart.NewCartRepo(db)
	orderRepo := order.NewOrderRepo(db)

	service := NewCartService(cartRepo, orderRepo)

	err = service.AddProduct(1, 1, 2)
	assert.NoError(t, err)

	var item models.CartItem
	err = db.Table("cart_items").First(&item, "cart_id = ? and product_id = ?", 1, 1).Error
	assert.NoError(t, err)
	assert.Equal(t, 2, item.Quantity)

	err = service.AddProduct(1, 1, 2)
	assert.NoError(t, err)

	item = models.CartItem{}
	err = db.First(&item, "cart_id = ? and product_id = ?", 1, 1).Error
	assert.NoError(t, err)
	assert.Equal(t, 4, item.Quantity)
}

func TestCartServiceImpl_GetCartItems(t *testing.T) {
	db := setupTestDB(t)
	repo := cart.NewCartRepo(db)
	orderRepo := order.NewOrderRepo(db)
	service := NewCartService(repo, orderRepo)

	err := db.Create(&models.Product{Name: "Product A", Price: 100}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Product{Name: "Product B", Price: 100}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Cart{UserID: 1}).Error
	assert.NoError(t, err)

	err = service.AddProduct(1, 1, 2)
	assert.NoError(t, err)
	err = service.AddProduct(1, 2, 2)
	assert.NoError(t, err)

	items, err := service.GetCartItems(1)
	assert.NoError(t, err)
	assert.Equal(t, items, &[]dto.ItemDTO{{Name: "Product A", Price: 100, Quantity: 2}, {Name: "Product B", Price: 100, Quantity: 2}})
}

func TestCartServiceImpl_ReduceProduct(t *testing.T) {
	db := setupTestDB(t)
	repo := cart.NewCartRepo(db)
	orderRepo := order.NewOrderRepo(db)

	service := NewCartService(repo, orderRepo)

	err := db.Create(&models.Product{Name: "Product A", Price: 100}).Error
	assert.NoError(t, err)

	err = service.AddProduct(1, 1, 10)
	assert.NoError(t, err)

	err = service.ReduceProduct(1, 1, 9)
	assert.NoError(t, err)

	var item models.CartItem
	err = db.First(&item, "cart_id = ? and product_id = ?", 1, 1).Error
	assert.NoError(t, err)
	assert.Equal(t, 1, item.Quantity)

	err = service.ReduceProduct(1, 1, 1)
	assert.NoError(t, err)

	err = db.First(&item, "cart_id = ? and product_id = ?", 1, 1).Error
	assert.Equal(t, err, gorm.ErrRecordNotFound)
}

func TestCartServiceImpl_RemoveProduct(t *testing.T) {
	db := setupTestDB(t)
	repo := cart.NewCartRepo(db)
	orderRepo := order.NewOrderRepo(db)
	service := NewCartService(repo, orderRepo)

	err := db.Create(&models.Product{Name: "Product A", Price: 100}).Error
	assert.NoError(t, err)
	err = service.AddProduct(1, 1, 10)
	assert.NoError(t, err)

	err = service.RemoveProduct(1, 1)
	assert.NoError(t, err)

	var item models.CartItem
	err = db.First(&item, "cart_id = ? and product_id = ?", 1, 1).Error
	assert.Equal(t, err, gorm.ErrRecordNotFound)
}

func TestCartServiceImpl_Checkout(t *testing.T) {
	db := setupTestDB(t)
	repo := cart.NewCartRepo(db)
	orderRepo := order.NewOrderRepo(db)
	service := NewCartService(repo, orderRepo)

	err := db.Create(&models.Cart{UserID: 1}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Product{Name: "Product A", Price: 100}).Error
	assert.NoError(t, err)
	err = service.AddProduct(1, 1, 10)
	assert.NoError(t, err)

	var item models.CartItem

	err = db.First(&item, "cart_id = ?", 1).Error
	assert.NoError(t, err)

	err = service.Checkout(1)
	assert.NoError(t, err)

	var newItem models.CartItem
	err = db.First(&newItem, "cart_id = ? and product_id = ?", 1, 1).Error
	assert.Equal(t, err, gorm.ErrRecordNotFound)

	var order models.Order
	err = db.Table("orders").Preload("Items").First(&order, "user_id = ?", 1).Error
	assert.NoError(t, err)
	assert.Equal(t, order.Items[0].ProductID, item.ProductID)
	assert.Equal(t, order.Items[0].Quantity, item.Quantity)

}
