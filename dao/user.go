package dao

import (
	"fmt"
	"srs_wrapper/database"
	. "srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/jameskeane/bcrypt"
)

func GetUserByID(id uint) *User {
	user := &User{}

	if err := database.DB.Preload("Group").First(user, id).Error; err != nil {
		fmt.Printf("GetUserByIdErr: %v\n", err)
	}

	return user
}

func GetUserByName(name string) *User {
	user := &User{Name: name}

	if err := database.DB.Preload("Group").Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserByUserNameErr: %v\n", err)
	}

	return user
}

func DeleteUserByID(id uint) {
	if err := database.DB.Delete(&User{}, id).Error; err != nil {
		fmt.Printf("DeleteUserByIdErr: %v\n", err)
	}
}

func GetAllUsers() []*User {
	return GetAllUsersWithParam("", "", "", 0, 0)
}

func GetAllUsersWithParam(name, displayName, orderBy string, offset, limit int) (users []*User) {
	user := &User{
		Name:        name,
		DisplayName: displayName,
	}
	if err := database.DB.Preload("Group").Where(user).Find(&users).Error; err != nil {
		fmt.Printf("GetAllUserErr: %v\n", err)
	}
	return
}

func CreateUser(ujson *UserJson) (*User, error) {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(ujson.Password, salt)
	if ujson.DisplayName == "" {
		ujson.DisplayName = ujson.Name
	}
	if ujson.GroupID == 0 {
		ujson.GroupID = GuestGroupID
	}

	user := &User{
		Name:        ujson.Name,
		Password:    string(hash),
		DisplayName: ujson.DisplayName,
		GroupID:     ujson.GroupID,
	}

	if err := database.DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr: %v\n", err)
		return nil, err
	}

	return user, nil
}

func UpdateUser(ujson *UserJson, id uint) (*User, error) {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(ujson.Password, salt)

	user := &User{
		Name:        ujson.Name,
		GroupID:     ujson.GroupID,
		DisplayName: ujson.DisplayName,
	}
	user.ID = id
	if ujson.Password != "" {
		user.Password = string(hash)
	}

	if err := database.DB.Model(&user).Updates(user).Error; err != nil {
		fmt.Printf("UpdateUserErr: %v\n", err)
		return nil, err
	}

	return user, nil
}

func CheckLogin(name, password string) (bool, string, error) {
	user := GetUserByName(name)
	if user.ID == 0 {
		return false, "", nil
	} else if ok := bcrypt.Match(password, user.Password); !ok {
		return false, "", nil
	} else if token, err := util.GetJwtString(user.ID); err != nil {
		return false, "", err
	} else {
		return true, token, nil
	}
}
