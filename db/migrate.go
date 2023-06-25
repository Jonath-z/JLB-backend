package db

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
)

func Migrate() {
	config.DB.AutoMigrate(&entities.UserEntity{})
	config.DB.AutoMigrate(&entities.ProductEntity{})
}
