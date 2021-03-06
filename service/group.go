package service

import (
	"errors"
	"fmt"
	"srs_wrapper/dao"
	"srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func GetGroupByID(id uint) *model.ApiJson {
	g, err := dao.GetGroupByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(g.ToJson(), "获取成功")
}

func GetGroupAllInfoByID(id uint) *model.ApiJson {
	g, err := dao.GetGroupByIDWithPerms(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(g.ToJson(), "获取成功")
}

func GetGroupByName(name string) *model.ApiJson {
	g, err := dao.GetGroupByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(g.ToJson(), "获取成功")
}

func GetGroupAllInfoByName(name string) *model.ApiJson {
	g, err := dao.GetGroupByNameWithPerms(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(g.ToJson(), "获取成功")
}

func CreateGroup(id uint, aul *model.GroupJson) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err.(validator.ValidationErrors))
		// for _, e := range err.(validator.ValidationErrors) {
		// 	fmt.Println()
		// 	fmt.Println(e.Namespace())
		// 	fmt.Println(e.Field())
		// 	fmt.Println(e.Type())
		// 	fmt.Println(e.Param())
		// 	fmt.Println()
		// }
	}
	user, err := dao.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	if !dao.HasPermission(user.GroupID, "admin.account") {
		return model.ErrorNoPermissions(fmt.Errorf("您没有权限组管理权限"))
	}
	g, err := dao.CreateGroup(aul, []*model.Permission{})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorInsertDatabase(err)
		}
	}
	return model.SuccessCreate(g.ToJson(), "创建成功")
}

func UpdateGroup(qid, id uint, aul *model.GroupJson) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err.(validator.ValidationErrors))
		// for _, e := range err.(validator.ValidationErrors) {
		// 	fmt.Println()
		// 	fmt.Println(e.Namespace())
		// 	fmt.Println(e.Field())
		// 	fmt.Println(e.Type())
		// 	fmt.Println(e.Param())
		// 	fmt.Println()
		// }
	}
	user, err := dao.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	if !dao.HasPermission(user.GroupID, "admin.account") {
		return model.ErrorNoPermissions(fmt.Errorf("您没有权限组管理权限"))
	}
	g, err := dao.UpdateGroup(aul, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorUpdateDatabase(err)
		}
	}
	return model.SuccessUpdate(g.ToJson(), "更新成功")
}

func DeleteGroup(qid, id uint) *model.ApiJson {
	user, err := dao.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	if !dao.HasPermission(user.GroupID, "admin.account") {
		return model.ErrorNoPermissions(fmt.Errorf("您没有权限组管理权限"))
	}
	err = dao.DeleteGroupByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorDeleteDatabase(err)
		}
	}
	return model.SuccessUpdate(nil, "删除成功")
}

func GetAllGroups(aul *model.AllGroupReq) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err)
	}
	if aul.Offset == 0 {
		aul.Offset = 1
	}
	if aul.Limit == 0 {
		aul.Limit = 20
	}
	groups, err := dao.GetAllGroupsWithParam(aul.Name, aul.OrderBy, aul.Offset, aul.Limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	gs := []*model.GroupJson{}
	for _, g := range groups {
		gs = append(gs, g.ToJson())
	}
	return model.Success(gs, "查询成功")
}
