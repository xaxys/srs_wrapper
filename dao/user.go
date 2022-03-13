package dao

import (
	"fmt"
	"srs_wrapper/database"
	. "srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/jameskeane/bcrypt"
)

func GetUserByID(id uint) (*User, error) {
	user := &User{}

	if err := database.DB.First(user, id).Error; err != nil {
		fmt.Printf("GetUserByIDErr: %v\n", err)
		return nil, err
	}

	return user, nil
}

func GetUserByIDWithGroup(id uint) (*User, error) {
	user := &User{}

	if err := database.DB.First(user, id).Error; err != nil {
		fmt.Printf("GetUserByIDWithGroupErr: %v\n", err)
		return nil, err
	}
	group, err := GetGroupByIDWithPerms(user.GroupID)
	if err != nil {
		fmt.Printf("GetUserByIDWithGroupErr: %v\n", err)
		return nil, err
	}
	user.Group = group

	return user, nil
}

func GetUserByName(name string) (*User, error) {
	user := &User{Name: name}

	if err := database.DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserByUserNameErr: %v\n", err)
		return nil, err
	}

	return user, nil
}

func GetUserByNameWithGroup(name string) (*User, error) {
	user := &User{Name: name}

	if err := database.DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserByUserNameErr: %v\n", err)
		return nil, err
	}
	group, err := GetGroupByIDWithPerms(user.GroupID)
	if err != nil {
		fmt.Printf("GetUserByNameWithGroupErr: %v\n", err)
		return nil, err
	}
	user.Group = group

	return user, nil
}

func DeleteUserByID(id uint) error {
	if err := database.DB.Delete(&User{}, id).Error; err != nil {
		fmt.Printf("DeleteUserByIdErr: %v\n", err)
		return err
	}
	return nil
}

func GetAllUsers() ([]*User, error) {
	return GetAllUsersWithParam("", "", "", 0, 0)
}

func GetAllUsersWithParam(name, displayName, orderBy string, offset, limit int) (users []*User, err error) {
	user := &User{
		Name:        name,
		DisplayName: displayName,
	}
	if err = database.DB.Where(user).Find(&users).Error; err != nil {
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
		ujson.GroupID = GetGuestGroupID()
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

func CheckLogin(name, password string) (string, error) {
	user, err := GetUserByName(name)
	if err != nil {
		return "", nil
	} else if ok := bcrypt.Match(password, user.Password); !ok {
		return "", nil
	} else if token, err := util.GetJwtString(user.ID); err != nil {
		return "", err
	} else {
		return token, nil
	}
}
