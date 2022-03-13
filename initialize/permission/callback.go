package permission

import (
	"errors"
	"fmt"
	"srs_wrapper/dao"
	. "srs_wrapper/model"

	"gorm.io/gorm"
)

func CreateDefaultCallbackPerm() {
	publish := &PermissionJson{
		Name:        "callback.publish",
		DisplayName: "推流权限",
		Description: "推流权限",
		Default:     false,
	}

	if _, err := dao.GetPermissionByName(publish.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Create Default Permission: %s\n", publish.Name)
			dao.CreatePermission(publish)
		} else {
			panic(err)
		}
	}

	play := &PermissionJson{
		Name:        "callback.play",
		DisplayName: "拉流权限",
		Description: "拉流权限",
		Default:     true,
	}

	if _, err := dao.GetPermissionByName(play.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Create Default Permission: %s\n", play.Name)
			dao.CreatePermission(play)
		} else {
			panic(err)
		}
	}
}
