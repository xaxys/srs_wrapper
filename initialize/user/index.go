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

	user, _ := dao.GetUserByID(1)

	if user.ID == 0 {
		fmt.Println("Create Default Administrator Account")
		u, err := dao.CreateUser(aul)
		if err != nil {
			panic(fmt.Errorf("Failed to create administrator account: %v\n", err))
		}
		return u
	} else {
		return user
	}
}
