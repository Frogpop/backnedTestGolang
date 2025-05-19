package postgres

import (
	"backnedTestGolang/internal/config"
	"backnedTestGolang/internal/models"
	"backnedTestGolang/pkg/database"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewPostgresDB(cfg config.StorageConfig) (*gorm.DB, error) {
	postgesURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SslMode,
	)
	if err := database.WaitForPostgres(postgesURL, 10, 5*time.Second); err != nil {

	}

	db, err := gorm.Open(postgres.Open(postgesURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
