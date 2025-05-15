package postgres

import (
	"backnedTestGolang/internal/config"
	"backnedTestGolang/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(postgesURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Product{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
