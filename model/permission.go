package model

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name        string `gorm:"unique; not null VARCHAR(191)"`
	DisplayName string `gorm:"VARCHAR(191)"`
	Description string `gorm:"VARCHAR(191)"`
	Default     bool   `gorm:"not null TINYINT(1)"`
}

type PermissionJson struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required,gte=4,lte=50"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Default     bool   `json:"default" validate:"required"`
}

type AllPermissionReq struct {
	Name    string `json:"name" validate:"gte=2,lte=50"`
	Default string `json:"default"`
	OrderBy string `json:"order_by"`
	Limit   int    `json:"limit" validate:"number"`
	Offset  int    `json:"offset" validate:"number"`
}

func (perm *Permission) ToJson() *PermissionJson {
	return &PermissionJson{
		ID:          perm.ID,
		Name:        perm.Name,
		DisplayName: perm.DisplayName,
		Description: perm.Description,
		Default:     perm.Default,
	}
}
