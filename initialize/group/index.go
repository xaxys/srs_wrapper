package group

import (
	"errors"
	"fmt"

	"srs_wrapper/dao"
	. "srs_wrapper/model"

	"gorm.io/gorm"
)

func CreateDefaultGroups() {
	CreateDefaultAdminGroup()
}

func CreateDefaultAdminGroup() {
	aul := &GroupJson{
		Name:        "admin",
		DisplayName: "超级管理员",
		Description: "超级管理员",
	}

	if _, err := dao.GetGroupByName(aul.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Create Default Group: %s\n", aul.Name)
			perms, _ := dao.GetAllPermissions()
			dao.CreateGroup(aul, perms)
		} else {
			panic(err)
		}
	}

	guest := &GroupJson{
		Name:        "guest",
		DisplayName: "访客",
		Description: "访客",
	}

	if _, err := dao.GetGroupByName(guest.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Create Default Group: %s\n", guest.Name)
			perms, _ := dao.GetAllPermissionsWithParam("", true, "", 0, 0)
			dao.CreateGroup(guest, perms)
		} else {
			panic(err)
		}
	}
}
