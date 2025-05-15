package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint64
	Items  []CartItem `gorm:"foreignkey:CartID"`
}

type CartItem struct {
	CartID    uint64 `gorm:"primaryKey;foreignKey:CartID"`
	ProductID uint64 `gorm:"primaryKey;foreignkey:ProductID"`
	Quantity  int
}
