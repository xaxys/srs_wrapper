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

func GetPermissionByID(id uint) *model.ApiJson {
	p, err := dao.GetPermissionByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(p.ToJson(), "获取成功")
}

func GetPermissionByName(name string) *model.ApiJson {
	p, err := dao.GetPermissionByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(p.ToJson(), "获取成功")
}

func CreatePermission(id uint, aul *model.PermissionJson) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err.(validator.ValidationErrors))
		// for _, e := ranpe err.(validator.ValidationErrors) {
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
		return model.ErrorNoPermissions(fmt.Errorf("您没有权限管理权限"))
	}
	p, err := dao.CreatePermission(aul)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorInsertDatabase(err)
		}
	}
	return model.SuccessCreate(p.ToJson(), "创建成功")
}

func UpdatePermission(qid, id uint, aul *model.PermissionJson) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err.(validator.ValidationErrors))
		// for _, e := ranpe err.(validator.ValidationErrors) {
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
		return model.ErrorNoPermissions(fmt.Errorf("您没有权限管理权限"))
	}
	p, err := dao.UpdatePermission(aul, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorUpdateDatabase(err)
		}
	}
	return model.SuccessUpdate(p.ToJson(), "更新成功")
}

func DeletePermission(qid, id uint) *model.ApiJson {
	user, err := dao.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	if !dao.HasPermission(user.GroupID, "admin.account") {
		return model.ErrorNoPermissions(fmt.Errorf("您没有权限管理权限"))
	}
	err = dao.DeletePermissionByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorDeleteDatabase(err)
		}
	}
	return model.SuccessUpdate(nil, "删除成功")
}

func GetAllPermissions(aul *model.AllPermissionReq) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err)
	}
	if aul.Offset == 0 {
		aul.Offset = 1
	}
	if aul.Limit == 0 {
		aul.Limit = 20
	}
	perms, err := dao.GetAllPermissionsWithParam(aul.Name, aul.Default == "true", aul.OrderBy, aul.Offset, aul.Limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	ps := []*model.PermissionJson{}
	for _, p := range perms {
		ps = append(ps, p.ToJson())
	}
	return model.Success(ps, "查询成功")
}
