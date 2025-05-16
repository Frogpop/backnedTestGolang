package cart

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
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

	if err := db.AutoMigrate(&models.Cart{}, &models.CartItem{}, &models.Product{}); err != nil {
		panic("migration failed")
	}

	return db
}

func TestCartRepo_CreateItem(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	err := repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})
	assert.NoError(t, err)

	var item models.CartItem
	err = db.Where("cart_id = ? AND product_id = ?", 1, 1).First(&item).Error
	assert.NoError(t, err)
	assert.Equal(t, models.CartItem{CartID: 1, ProductID: 1, Quantity: 10}, item)

	err = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})
	assert.Error(t, err)
}

func TestCartRepo_DeleteItem(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	_ = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})

	err := repo.RemoveItem(1, 1)
	assert.NoError(t, err)

	var item models.CartItem
	err = db.Where("cart_id = ? AND product_id = ?", 1, 1).First(&item).Error
	assert.Error(t, err)
}

func TestCartRepo_UpdateItem(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})

	err := repo.UpdateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 20})
	assert.NoError(t, err)

	var item models.CartItem
	err = db.Where("cart_id = ? AND product_id = ?", 1, 1).First(&item).Error
	assert.NoError(t, err)
	assert.Equal(t, models.CartItem{CartID: 1, ProductID: 1, Quantity: 20}, item)

}

func TestCartRepo_GetCartItem(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	err := db.Create(&models.Product{Name: "Product 1", Price: 100}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Product{Name: "Product 2", Price: 200}).Error
	assert.NoError(t, err)

	err = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})
	assert.NoError(t, err)
	err = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 2, Quantity: 20})
	assert.NoError(t, err)

	items, err := repo.GetCartItems(1)
	assert.NoError(t, err)
	assert.NotZero(t, items)

	var rows *[]dto.ItemDTO
	err = db.Table("cart_items").Select("name", "price", "quantity").Joins("join products on products.id = cart_items.product_id").Where("cart_id = ?", 1).Scan(&rows).Error
	assert.NoError(t, err)
	assert.Equal(t, items, rows)
}

func TestCartRepo_ClearCart(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	db.Create(&models.Cart{UserID: 1})

	err := db.Create(&models.Product{Name: "Product 1", Price: 100}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Product{Name: "Product 2", Price: 200}).Error
	assert.NoError(t, err)

	err = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})
	assert.NoError(t, err)
	err = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 2, Quantity: 20})
	assert.NoError(t, err)

	err = repo.ClearCart(1)
	assert.NoError(t, err)

	var cart models.Cart

	err = db.Table("cart_items").Where("cart_id = ?", 1).First(&cart).Error
	assert.Error(t, err)
	assert.Equal(t, models.Cart{UserID: 0}, cart)
}

func TestCartRepo_GetCartItems(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	err := db.Create(&models.Cart{UserID: 1}).Error
	assert.NoError(t, err)

	cart, err := repo.GetCart(1)
	assert.NoError(t, err)
	assert.NotZero(t, cart)
	assert.Equal(t, cart.UserID, uint64(1))
}

func TestCartRepo_CheckItem(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepo(db)

	err := db.Create(&models.Cart{UserID: 1}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Product{Name: "Product 1", Price: 100}).Error

	err = repo.CreateItem(&models.CartItem{CartID: 1, ProductID: 1, Quantity: 10})
	assert.NoError(t, err)

	item, err := repo.CheckItem(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, &models.CartItem{CartID: 1, ProductID: 1, Quantity: 10}, item)
}
