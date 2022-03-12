package dao

import (
	"fmt"
	"srs_wrapper/database"
	. "srs_wrapper/model"
)

func GetGroupByID(id uint) (*Group, error) {
	group := &Group{}

	if err := database.DB.First(group, id).Error; err != nil {
		fmt.Printf("GetGroupByIdErr: %v\n", err)
		return nil, err
	}

	return group, nil
}

func GetGroupByIDWithPerms(id uint) (*Group, error) {
	group := &Group{}

	if err := database.DB.Preload("Perms").First(group, id).Error; err != nil {
		fmt.Printf("GetGroupByIdErr: %v\n", err)
		return nil, err
	}

	return group, nil
}

func GetGroupByName(name string) (*Group, error) {
	group := &Group{Name: name}

	if err := database.DB.Where(group).First(group).Error; err != nil {
		fmt.Printf("GetGroupByNameErr: %v\n", err)
		return nil, err
	}

	return group, nil
}

func GetGroupByNameWithPerms(name string) (*Group, error) {
	group := &Group{Name: name}

	if err := database.DB.Preload("Perms").Where(group).First(group).Error; err != nil {
		fmt.Printf("GetGroupByNameErr: %v\n", err)
		return nil, err
	}

	return group, nil
}

func DeleteGroupByID(id uint) error {
	if err := database.DB.Delete(&Group{}, id).Error; err != nil {
		fmt.Printf("DeleteGroupErr: %v\n", err)
		return err
	}

	return nil
}

func GetAllGroups() ([]*Group, error) {
	return GetAllGroupsWithParam("", "", 0, 0)
}

func GetAllGroupsWithParam(name, orderBy string, offset, limit int) (groups []*Group, err error) {
	group := &Group{Name: name}

	if err = database.DB.Where(group).Find(&groups).Error; err != nil {
		fmt.Printf("GetAllGroupErr: %v\n", err)
	}
	return
}

func CreateGroup(gjson *GroupJson, perms []*Permission) (*Group, error) {
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
		return nil, err
	}

	if err := database.DB.Model(&group).Association("Perms").Append(perms); err != nil {
		fmt.Printf("AppendPermsErr: %v\n", err)
		return nil, err
	}

	return group, nil
}

func UpdateGroup(gjson *GroupJson, id uint) (*Group, error) {
	group := &Group{
		Name:        gjson.Name,
		DisplayName: gjson.DisplayName,
		Description: gjson.Description,
	}
	group.ID = id

	if err := database.DB.Model(&group).Updates(group).Error; err != nil {
		fmt.Printf("UpdateGroupErr: %v\n", err)
		return nil, err
	}

	return group, nil
}

func HasPermission(id uint, perm string) bool {
	group := &Group{}
	group.ID = id
	perms := []*Permission{}
	if err := database.DB.Model(&group).Association("Perms").Find(&perms, "name = ?", perm); err != nil {
		fmt.Printf("Database HasPermission Error: %v\n", err)
	}
	return len(perms) > 0
}
