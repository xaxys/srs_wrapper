package dao

import (
	"fmt"
	"srs_wrapper/database"
	. "srs_wrapper/model"
)

func GetPermissionByID(id uint) *Permission {
	permission := &Permission{}

	if err := database.DB.First(permission, id).Error; err != nil {
		fmt.Printf("GetPermissionByIdError: %v\n", err)
	}

	return permission
}

func GetPermissionByName(name string) *Permission {
	permission := &Permission{Name: name}

	if err := database.DB.Where(permission).First(permission).Error; err != nil {
		fmt.Printf("GetPermissionByNameError: %v\n", err)
	}

	return permission
}

func DeletePermissionByID(id uint) {
	if err := database.DB.Delete(&Permission{}, id).Error; err != nil {
		fmt.Printf("DeletePermissionByIdError: %v\n", err)
	}
}

func GetAllPermissions() (permissions []*Permission) {
	if err := database.DB.Find(&permissions).Error; err != nil {
		fmt.Printf("GetAllPermissionsError: %v\n", err)
	}
	return
}

func GetAllPermissionsWithParam(name string, dft bool, orderBy string, offset, limit int) (permissions []*Permission) {
	permission := &Permission{
		Name:    name,
		Default: dft,
	}

	if err := database.DB.Where(permission).Find(&permissions).Error; err != nil {
		fmt.Printf("GetAllPermissionsError: %v\n", err)
	}

	return
}

func CreatePermission(aul *PermissionJson) *Permission {
	permission := &Permission{
		Name:        aul.Name,
		DisplayName: aul.DisplayName,
		Description: aul.Description,
		Default:     aul.Default,
	}

	if err := database.DB.Create(permission).Error; err != nil {
		fmt.Printf("CreatePermissionError: %v\n", err)
	}

	return permission
}

func UpdatePermission(pjson *PermissionJson, id uint) *Permission {
	permission := &Permission{
		Name:        pjson.Name,
		DisplayName: pjson.DisplayName,
		Description: pjson.Description,
		Default:     pjson.Default,
	}
	permission.ID = id

	if err := database.DB.Model(&permission).Updates(pjson).Error; err != nil {
		fmt.Printf("UpdatePermissionError: %v\n", err)
	}

	return permission
}
