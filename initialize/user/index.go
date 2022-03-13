package group

import (
	"errors"
	"fmt"

	"srs_wrapper/config"
	"srs_wrapper/dao"
	. "srs_wrapper/model"

	"gorm.io/gorm"
)

func CreateDefaultUsers() {
	CreateSystemAdmin()
}

func CreateSystemAdmin() {
	aul := &UserJson{
		Name:        config.AppConfig.GetString("admin.name"),
		DisplayName: config.AppConfig.GetString("admin.display_name"),
		Password:    config.AppConfig.GetString("admin.password"),
		GroupID:     1,
	}

	if _, err := dao.GetUserByID(1); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Create Default Administrator Account")
			dao.CreateUser(aul)
		} else {
			panic(err)
		}
	}
}
