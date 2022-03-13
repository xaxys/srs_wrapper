package permission

import (
	"errors"
	"fmt"
	"srs_wrapper/dao"
	. "srs_wrapper/model"

	"gorm.io/gorm"
)

func CreateDefaultPermissions() {
	CreateDefaultAdminPerm()
	CreateDefaultCallbackPerm()
}

func CreateDefaultAdminPerm() {
	aul := &PermissionJson{
		Name:        "admin.account",
		DisplayName: "账号管理权限",
		Description: "账号管理权限",
		Default:     false,
	}

	if _, err := dao.GetPermissionByName(aul.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Create Default Permission: %s\n", aul.Name)
			dao.CreatePermission(aul)
		} else {
			panic(err)
		}
	}
}
