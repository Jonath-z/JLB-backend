package entities

import "gorm.io/gorm"

type ProductEntity struct {
	gorm.Model
	ProductId string
	Name      string
	Price     float64
	Image     string
	UserId    uint `gorm:"column:user_id"`
}
