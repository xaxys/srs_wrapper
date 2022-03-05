package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"unique; VARCHAR(191)"`
	Password    string `gorm:"not null VARCHAR(191)"`
	DisplayName string `gorm:"not null VARCHAR(191)"`
	GroupID     uint
	Group       Group
}

type UserJson struct {
	Name        string `json:"name" validate:"required,gte=2,lte=50"`
	Password    string `json:"password" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,gte=2,lte=50"`
	GroupID     uint   `json:"Group_id" validate:"required"`
}

type AllUserReq struct {
	Name        string `json:"name" validate:"gte=2,lte=50"`
	DisplayName string `json:"display_name" validate:"gte=2,lte=50"`
	OrderBy     string `json:"order_by"`
	Limit       int    `json:"limit" validate:"number"`
	Offset      int    `json:"offset" validate:"number"`
}
