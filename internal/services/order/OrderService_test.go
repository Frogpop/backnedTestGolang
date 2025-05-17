package order

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/models"
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

func TestOrderServiceImpl_GetOrders(t *testing.T) {
	db := setupTestDB(t)
	repo := order.NewOrderRepo(db)

	service := NewOrderService(repo)

	err := db.Create(&models.Product{Name: "ProductA", Price: 100.0}).Error
	assert.NoError(t, err)
	err = db.Create(&models.Product{Name: "ProductB", Price: 100.0}).Error
	assert.NoError(t, err)

	err = db.Create(&models.Order{UserID: 1, Items: []models.OrderItem{{ProductID: 1, Quantity: 1}, {ProductID: 2, Quantity: 1}}}).Error
	assert.NoError(t, err)

	orders, err := service.GetUserOrders(1)
	assert.NoError(t, err)

	var q_orders []dto.ItemDTO
	err = db.Table("order_items").Select("name, price, quantity").Joins("join products p on p.id = order_items.product_id").Where("order_id = ?", 1).Find(&q_orders).Error
	assert.NoError(t, err)
	assert.NotZero(t, q_orders)
	assert.Equal(t, orders.Orders[0].Items, q_orders)
}

func TestOrderServiceImpl_ChangeOrderStatus(t *testing.T) {
	db := setupTestDB(t)
	repo := order.NewOrderRepo(db)
	service := NewOrderService(repo)

	err := db.Create(&models.Order{UserID: 1, Status: "created"}).Error
	assert.NoError(t, err)

	err = service.ChangeOrderStatus(1, "changed")
	assert.NoError(t, err)

	var order models.Order
	err = db.Where("id = ?", 1).First(&order).Error
	assert.NoError(t, err)

	assert.Equal(t, order.Status, "changed")

}
