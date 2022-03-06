package dao

import (
	"fmt"
	"srs_wrapper/database"
	. "srs_wrapper/model"
)

func GetGroupByID(id uint) *Group {
	group := &Group{}

	if err := database.DB.Preload("Perms").First(group, id).Error; err != nil {
		fmt.Printf("GetGroupByIdErr: %v\n", err)
	}

	return group
}

func GetGroupByName(name string) *Group {
	group := &Group{Name: name}

	if err := database.DB.Preload("Perms").Where(group).First(group).Error; err != nil {
		fmt.Printf("GetGroupByNameErr: %v\n", err)
	}

	return group
}

func DeleteGroupByID(id uint) {
	if err := database.DB.Delete(&Group{}, id).Error; err != nil {
		fmt.Printf("DeleteGroupErr: %v\n", err)
	}
}

func GetAllGroups() []*Group {
	return GetAllGroupsWithParam("", "", 0, 0)
}

func GetAllGroupsWithParam(name, orderBy string, offset, limit int) (groups []*Group) {
	group := &Group{Name: name}

	if err := database.DB.Preload("Perms").Where(group).Find(&groups).Error; err != nil {
		fmt.Printf("GetAllGroupErr: %v\n", err)
	}
	return
}

func CreateGroup(gjson *GroupJson, perms []*Permission) *Group {
	group := &Group{
		Name:        gjson.Name,
		DisplayName: gjson.DisplayName,
		Description: gjson.Description,
	}

	if gjson.DisplayName == "" {
		gjson.DisplayName = gjson.Name
	}

	if err := database.DB.Create(group).Error; err != nil {
		fmt.Printf("CreateGroupErr: %v\n", err)
	}

	if err := database.DB.Model(&group).Association("Perms").Append(perms); err != nil {
		fmt.Printf("AppendPermsErr: %v\n", err)
	}

	return group
}

func UpdateGroup(gjson *GroupJson, id uint) *Group {
	group := &Group{
		Name:        gjson.Name,
		DisplayName: gjson.DisplayName,
		Description: gjson.Description,
	}
	group.ID = id

	if err := database.DB.Model(&group).Updates(group).Error; err != nil {
		fmt.Printf("UpdatGroupErr: %v\n", err)
	}

	return group
}

func HasPermission(group *Group, perm string) bool {
	perms := []*Permission{}
	if err := database.DB.Model(&group).Association("Perms").Find(&perms, "name = ?", perm); err != nil {
		fmt.Printf("Database HasPermission Error: %v\n", err)
	}
	return len(perms) > 0
}
