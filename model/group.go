package model

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string        `gorm:"unique; not null VARCHAR(191)"`
	DisplayName string        `gorm:"VARCHAR(191)"`
	Description string        `gorm:"VARCHAR(191)"`
	Perms       []*Permission `gorm:"many2many:group_perms;"`
}

type GroupJson struct {
	Name        string `json:"name" validate:"required,gte=4,lte=50"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

type AllGroupReq struct {
	Name    string `json:"name" validate:"gte=2,lte=50"`
	OrderBy string `json:"order_by"`
	Limit   int    `json:"limit" validate:"number"`
	Offset  int    `json:"offset" validate:"number"`
}

func (g *Group) HasPermission(perm string) bool {
	for _, p := range g.Perms {
		if p.Name == perm {
			return true
		}
	}
	return false
}
