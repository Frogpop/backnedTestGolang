package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     uint64
	UserID uint64
	Status string
	Items  []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	OrderID   uint64 `gorm:"primaryKey;foreignKey:OrderID"`
	ProductID uint64 `gorm:"primaryKey:foreignKey:ProductID"`
	Quantity  int
	//Price     float64
}
