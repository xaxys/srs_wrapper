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
	ID          uint              `json:"id"`
	Name        string            `json:"name" validate:"required,gte=4,lte=50"`
	DisplayName string            `json:"display_name"`
	Description string            `json:"description"`
	Perms       []*PermissionJson `json:"perms,omitempty"`
}

type AllGroupReq struct {
	Name    string `json:"name" validate:"gte=2,lte=50"`
	OrderBy string `json:"order_by"`
	Limit   int    `json:"limit" validate:"number"`
	Offset  int    `json:"offset" validate:"number"`
}

func (perm *Group) ToJson() *GroupJson {
	perms := []*PermissionJson{}
	for _, p := range perm.Perms {
		perms = append(perms, p.ToJson())
	}
	return &GroupJson{
		ID:          perm.ID,
		Name:        perm.Name,
		DisplayName: perm.DisplayName,
		Description: perm.Description,
		Perms:       perms,
	}
}
