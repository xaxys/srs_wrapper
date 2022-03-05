package controller

import (
	"fmt"
	"srs_wrapper/dao"
	"srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/kataras/iris/v12"
)

func GetProfile(ctx iris.Context) {
	userID := ctx.Values().Get("user_id").(uint)
	user := dao.GetUserByID(userID)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(model.ApiResponse(true, user, "操作成功"))
}

func GetUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	user := dao.GetUserByID(id)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(model.ApiResponse(true, user, "操作成功"))
}

func CreateUser(ctx iris.Context) {
	aul := &model.UserJson{}
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(err))
	} else if err := util.Validator.Struct(aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
		// for _, e := range err.(validator.ValidationErrors) {
		// 	fmt.Println()
		// 	fmt.Println(e.Namespace())
		// 	fmt.Println(e.Field())
		// 	fmt.Println(e.Type())
		// 	fmt.Println(e.Param())
		// 	fmt.Println()
		// }
	} else {
		u := dao.CreateUser(aul)
		ctx.StatusCode(iris.StatusOK)
		if u.ID == 0 {
			ctx.JSON(model.ApiResponse(false, u, "操作失败"))
		} else {
			ctx.JSON(model.ApiResponse(true, u, "操作成功"))
		}
	}
}

func UpdateUser(ctx iris.Context) {
	aul := &model.UserJson{}
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(err))
	} else if err := util.Validator.Struct(aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println()
		// 	fmt.Println(err.Namespace())
		// 	fmt.Println(err.Field())
		// 	fmt.Println(err.Type())
		// 	fmt.Println(err.Param())
		// 	fmt.Println()
		// }
	} else if id, err := ctx.Params().GetUint("id"); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
	} else {
		u := dao.UpdateUser(aul, id)
		ctx.StatusCode(iris.StatusOK)
		if u.ID == 0 {
			ctx.JSON(model.ApiResponse(false, u, "操作失败"))
		} else {
			ctx.JSON(model.ApiResponse(true, u, "操作成功"))
		}
	}

}

func DeleteUser(ctx iris.Context) {
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
	} else {
		dao.DeleteUserByID(id)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ApiResponse(true, nil, "删除成功"))
	}
}

func GetAllUsers(ctx iris.Context) {
	aul := &model.AllUserReq{}
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(err))
	} else if err := util.Validator.Struct(aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
	} else {
		if aul.Offset == 0 {
			aul.Offset = 1
		}
		if aul.Limit == 0 {
			aul.Limit = 20
		}
		users := dao.GetAllUsersWithParam(aul.Name, aul.DisplayName, aul.OrderBy, aul.Offset, aul.Limit)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ApiResponse(true, users, "操作成功"))
	}
}

func UserLogin(ctx iris.Context) {
	aul := &model.UserJson{}

	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(fmt.Errorf("请求参数错误")))
	} else {
		if UserNameErr := util.Validator.Var(aul.Name, "required,min=4,max=20"); UserNameErr != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(model.ErrorVerification(fmt.Errorf("用户名格式错误")))
		} else if PwdErr := util.Validator.Var(aul.Password, "required,min=5,max=20"); PwdErr != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(model.ErrorVerification(fmt.Errorf("密码格式错误")))
		} else {
			ok, token, err := dao.CheckLogin(aul.Name, aul.Password)
			if ok {
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(model.ApiResponse(ok, token, "登陆成功"))
			} else if err == nil {
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(model.ErrorUnauthorized(fmt.Errorf("用户名或密码错误")))
			} else {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.JSON(model.ErrorBuildJWT(err))
			}
		}
	}
}

func UserLogout(ctx iris.Context) {
	// aui := ctx.Values().GetString("user_id")
	// uid := uint(tools.Tool.ParseInt(aui, 0))
	// model.UserAdminLogout(uid)

	// ctx.StatusCode(http.StatusOK)
	// ctx.JSON(model.ApiResponse(true, nil, "退出登陆"))
	ctx.StatusCode(iris.StatusNotImplemented)
	ctx.JSON(model.ErrorInternalServer(fmt.Errorf("Not Implemented")))
}
