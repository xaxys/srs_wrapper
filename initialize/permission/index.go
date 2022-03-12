package permission

import (
	"fmt"
	"srs_wrapper/dao"
	. "srs_wrapper/model"
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

	if perm, _ := dao.GetPermissionByName(aul.Name); perm.ID == 0 {
		fmt.Printf("Create Default Permission: %s\n", aul.Name)
		dao.CreatePermission(aul)
	}
}
