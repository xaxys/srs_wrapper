package permission

import (
	"fmt"
	"srs_wrapper/dao"
	. "srs_wrapper/model"
)

func CreateDefaultCallbackPerm() {
	publish := &PermissionJson{
		Name:        "callback.publish",
		DisplayName: "推流权限",
		Description: "推流权限",
		Default:     false,
	}

	if perm, _ := dao.GetPermissionByName(publish.Name); perm.ID == 0 {
		fmt.Printf("Create Default Permission: %s\n", publish.Name)
		dao.CreatePermission(publish)
	}

	play := &PermissionJson{
		Name:        "callback.play",
		DisplayName: "拉流权限",
		Description: "拉流权限",
		Default:     true,
	}

	if perm, _ := dao.GetPermissionByName(play.Name); perm.ID == 0 {
		fmt.Printf("Create Default Permission: %s\n", play.Name)
		dao.CreatePermission(play)
	}
}
