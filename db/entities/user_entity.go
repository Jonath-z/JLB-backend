package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name             *string
	Email            string
	ProfileUrl       *string
	ProfileThumbnail *string
	IsVerified       bool
	Password         string
	UserId           string
}
