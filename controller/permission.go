package controller

import (
	"fmt"
	"srs_wrapper/dao"
	"srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/kataras/iris/v12"
)

func GetPermission(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	permission := dao.GetPermissionByID(id)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(model.ApiResponse(true, permission, "操作成功"))
}

func CreatePermission(ctx iris.Context) {
	perm := &model.PermissionJson{}
	if err := ctx.ReadJSON(&perm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(err))
	} else if err := util.Validator.Struct(perm); err != nil {
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
	} else if uid := ctx.Values().Get("user_id").(uint); !dao.HasPermission(&dao.GetUserByID(uid).Group, "admin.account") {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ErrorInsufficientPermissions(fmt.Errorf("您没有权限组管理权限")))
	} else {
		u := dao.CreatePermission(perm)
		ctx.StatusCode(iris.StatusOK)
		if u.ID == 0 {
			ctx.JSON(model.ApiResponse(false, u, "操作失败"))
		} else {
			ctx.JSON(model.ApiResponse(true, u, "操作成功"))
		}
	}
}

func UpdatePermission(ctx iris.Context) {
	perm := &model.PermissionJson{}
	if err := ctx.ReadJSON(&perm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(err))
	} else if err := util.Validator.Struct(perm); err != nil {
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
	} else if uid := ctx.Values().Get("user_id").(uint); !dao.HasPermission(&dao.GetUserByID(uid).Group, "admin.account") {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ErrorInsufficientPermissions(fmt.Errorf("您没有权限管理权限")))
	} else {
		u := dao.UpdatePermission(perm, id)
		ctx.StatusCode(iris.StatusOK)
		if u.ID == 0 {
			ctx.JSON(model.ApiResponse(false, u, "操作失败"))
		} else {
			ctx.JSON(model.ApiResponse(true, u, "操作成功"))
		}
	}

}

func DeletePermission(ctx iris.Context) {
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
	} else if uid := ctx.Values().Get("user_id").(uint); !dao.HasPermission(&dao.GetUserByID(uid).Group, "admin.account") {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ErrorInsufficientPermissions(fmt.Errorf("您没有权限管理权限")))
	} else {
		dao.DeletePermissionByID(id)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ApiResponse(true, nil, "删除成功"))
	}
}

func GetAllPermissions(ctx iris.Context) {
	// perm := &model.AllPermissionReq{}
	// if err := ctx.ReadJSON(&perm); err != nil {
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	ctx.JSON(model.ErrorIncompleteData(err))
	// } else if err := util.Validator.Struct(perm); err != nil {
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	ctx.JSON(model.ErrorVerification(err))
	// } else {
	// 	if perm.Offset == 0 {
	// 		perm.Offset = 1
	// 	}
	// 	if perm.Limit == 0 {
	// 		perm.Limit = 20
	// 	}
	// 	if perm.Default == "" {
	// 		permissions := dao.GetAllPermissions(perm.Name, perm.OrderBy, perm.Offset, perm.Limit)
	// 	}
	// 	permissions := dao.GetAllPermissionsWithParam(perm.Name, perm.OrderBy, perm.Offset, perm.Limit)

	// 	ctx.StatusCode(iris.StatusOK)
	// 	ctx.JSON(model.ApiResponse(true, permissions, "操作成功"))
	// }
	ctx.StatusCode(iris.StatusNotImplemented)
	ctx.JSON(model.ErrorInternalServer(fmt.Errorf("Not Implemented")))
}
