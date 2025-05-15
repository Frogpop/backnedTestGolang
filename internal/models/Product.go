package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	OrderItems []OrderItem `gorm:"foreignkey:ProductID"`
	CartItems  []CartItem  `gorm:"foreignkey:ProductID"`
}
