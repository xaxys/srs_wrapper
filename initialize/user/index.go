package group

import (
	"fmt"

	"srs_wrapper/config"
	"srs_wrapper/dao"
	. "srs_wrapper/model"
)

func CreateDefaultUsers() {
	CreateSystemAdmin()
}

func CreateSystemAdmin() *User {
	aul := &UserJson{
		Name:        config.AppConfig.GetString("admin.name"),
		DisplayName: config.AppConfig.GetString("admin.display_name"),
		Password:    config.AppConfig.GetString("admin.password"),
		GroupID:     1,
	}

	user := dao.GetUserByID(1)

	if user.ID == 0 {
		fmt.Println("Create Default Administrator Account")
		return dao.CreateUser(aul)
	} else {
		return user
	}
}