package order

import (
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

	if err := db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{}); err != nil {
		panic("migration failed")
	}

	return db
}

func TestOrderRepo_CreateOrder(t *testing.T) {
	db := setupTestDB(t)
	repo := NewOrderRepo(db)

	order := &models.Order{UserID: 1, Status: "created"}

	_, err := repo.CreateOrder(order)
	assert.NoError(t, err)
	assert.NotZero(t, order.ID)

	_, err = repo.CreateOrder(order)
	assert.Error(t, err)

}

func TestOrderRepo_UpdateOrder(t *testing.T) {
	db := setupTestDB(t)
	repo := NewOrderRepo(db)

	err := db.Create(&models.Order{UserID: 1, Status: "created"}).Error

	assert.NoError(t, err)

	err = repo.ChangeOrderStatus(1, "delivered")
	assert.NoError(t, err)

	order := &models.Order{}
	err = db.Table("orders").Where("id = ?", 1).First(order).Error
	assert.NoError(t, err)

	assert.Equal(t, "delivered", order.Status)
}
