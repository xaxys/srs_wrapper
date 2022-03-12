package controller

import (
	"srs_wrapper/model"
	"srs_wrapper/service"

	"github.com/kataras/iris/v12"
)

func GetPermission(ctx iris.Context) {
	id := ctx.Values().Get("user_id").(uint)
	response := service.GetPermissionByID(id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func CreatePermission(ctx iris.Context) {
	perm := &model.PermissionJson{}
	if err := ctx.ReadJSON(&perm); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	id := ctx.Values().Get("user_id").(uint)
	response := service.CreatePermission(id, perm)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func UpdatePermission(ctx iris.Context) {
	perm := &model.PermissionJson{}
	if err := ctx.ReadJSON(&perm); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	qid, _ := ctx.Params().GetUint("id")
	id := ctx.Values().Get("user_id").(uint)
	response := service.UpdatePermission(qid, id, perm)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func DeletePermission(ctx iris.Context) {
	qid, _ := ctx.Params().GetUint("id")
	id := ctx.Values().Get("user_id").(uint)
	response := service.DeletePermission(qid, id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func GetAllPermissions(ctx iris.Context) {
	req := &model.AllPermissionReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	response := service.GetAllPermissions(req)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}
