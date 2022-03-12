package group

import (
	"fmt"

	"srs_wrapper/dao"
	. "srs_wrapper/model"
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

	if group, _ := dao.GetGroupByName(aul.Name); group.ID == 0 {
		fmt.Printf("Create Default Group: %s\n", aul.Name)
		perms := dao.GetAllPermissions()
		dao.CreateGroup(aul, perms)
	}

	gst := &GroupJson{
		Name:        "guest",
		DisplayName: "访客",
		Description: "访客",
	}

	if group, _ := dao.GetGroupByName(gst.Name); group.ID == 0 {
		fmt.Printf("Create Default Group: %s\n", gst.Name)
		perms := dao.GetAllPermissionsWithParam("", true, "", 0, 0)
		dao.CreateGroup(gst, perms)
	}
}
