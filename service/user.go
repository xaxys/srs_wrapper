package service

import (
	"errors"
	"fmt"
	"srs_wrapper/dao"
	"srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func GetUserByID(id uint) *model.ApiJson {
	user, err := dao.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(user.ToJson(), "获取成功")
}

func GetUserAllInfoByID(id uint) *model.ApiJson {
	user, err := dao.GetUserByIDWithGroup(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(user.ToJson(), "获取成功")
}

func GetUserByName(name string) *model.ApiJson {
	user, err := dao.GetUserByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(user.ToJson(), "获取成功")
}

func GetUserAllInfoByName(name string) *model.ApiJson {
	user, err := dao.GetUserByNameWithGroup(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	return model.Success(user.ToJson(), "获取成功")
}

func CreateUser(id uint, aul *model.UserJson) *model.ApiJson {
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
	if aul.GroupID != 0 && !dao.HasPermission(user.GroupID, "admin.account") {
		return model.ErrorNoPermissions(fmt.Errorf("您没有指定权限组的权限"))
	}
	u, err := dao.CreateUser(aul)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorInsertDatabase(err)
		}
	}
	u.Password = ""
	return model.SuccessCreate(u.ToJson(), "创建成功")

}

func UpdateUser(qid, id uint, aul *model.UserJson) *model.ApiJson {
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
	authorized := dao.HasPermission(user.GroupID, "admin.account")
	if qid != id && !authorized {
		return model.ErrorNoPermissions(fmt.Errorf("您没有账号管理权限"))
	}
	if aul.GroupID != 0 && !authorized {
		return model.ErrorNoPermissions(fmt.Errorf("您没有修改权限组的权限"))
	}
	u, err := dao.UpdateUser(aul, qid)
	if err != nil {
		return model.ErrorInsertDatabase(err)
	}
	u.Password = ""
	return model.SuccessUpdate(u.ToJson(), "更新成功")
}

func DeleteUser(qid, id uint) *model.ApiJson {
	user, err := dao.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	if qid != id && !dao.HasPermission(user.GroupID, "admin.account") {
		return model.ErrorNoPermissions(fmt.Errorf("您没有账号管理权限"))
	}
	err = dao.DeleteUserByID(qid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorDeleteDatabase(err)
		}
	}
	return model.SuccessUpdate(nil, "删除成功")
}

func GetAllUsers(aul *model.AllUserReq) *model.ApiJson {
	if err := util.Validator.Struct(aul); err != nil {
		return model.ErrorVerification(err)
	}
	if aul.Offset == 0 {
		aul.Offset = 1
	}
	if aul.Limit == 0 {
		aul.Limit = 20
	}
	users, err := dao.GetAllUsersWithParam(aul.Name, aul.DisplayName, aul.OrderBy, aul.Offset, aul.Limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrorNotFound(err)
		} else {
			return model.ErrorQueryDatabase(err)
		}
	}
	us := []*model.UserJson{}
	for _, u := range users {
		us = append(us, u.ToJson())
	}
	return model.Success(us, "操作成功")
}

func UserLogin(ctx iris.Context) *model.ApiJson {
	aul := &model.UserJson{}

	if err := ctx.ReadJSON(&aul); err != nil {
		return model.ErrorInvalidData(fmt.Errorf("请求参数错误"))
	}
	if UserNameErr := util.Validator.Var(aul.Name, "required,min=4,max=20"); UserNameErr != nil {
		return model.ErrorVerification(fmt.Errorf("用户名格式错误"))
	}
	if PwdErr := util.Validator.Var(aul.Password, "required,min=5,max=20"); PwdErr != nil {
		return model.ErrorVerification(fmt.Errorf("密码格式错误"))
	}
	token, err := dao.CheckLogin(aul.Name, aul.Password)
	if token != "" {
		return model.Success(token, "登陆成功")
	}
	if err != nil {
		return model.ErrorBuildJWT(err)
	}
	return model.ErrorUnauthorized(fmt.Errorf("用户名或密码错误"))
}

func UserLogout(ctx iris.Context) *model.ApiJson {
	// aui := ctx.Values().GetString("user_id")
	// uid := uint(tools.Tool.ParseInt(aui, 0))
	// model.UserAdminLogout(uid)

	// ctx.StatusCode(http.StatusOK)
	// return model.Success(true, nil, "退出登陆"))
	return model.ErrorInternalServer(fmt.Errorf("Not Implemented"))
}
