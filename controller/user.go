package controller

import (
	"srs_wrapper/model"
	"srs_wrapper/service"

	"github.com/kataras/iris/v12"
)

func GetProfile(ctx iris.Context) {
	id := ctx.Values().Get("user_id").(uint)
	response := service.GetUser(id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func GetUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	response := service.GetUser(id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func CreateUser(ctx iris.Context) {
	aul := &model.UserJson{}
	if err := ctx.ReadJSON(&aul); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	id := ctx.Values().Get("user_id").(uint)
	response := service.CreateUser(id, aul)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func UpdateUser(ctx iris.Context) {
	aul := &model.UserJson{}
	if err := ctx.ReadJSON(&aul); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	qid, _ := ctx.Params().GetUint("id")
	id := ctx.Values().Get("user_id").(uint)
	response := service.UpdateUser(qid, id, aul)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func DeleteUser(ctx iris.Context) {
	qid, _ := ctx.Params().GetUint("id")
	id := ctx.Values().Get("user_id").(uint)
	response := service.DeleteUser(qid, id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func GetAllUsers(ctx iris.Context) {
	req := &model.AllUserReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	response := service.GetAllUsers(req)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func UserLogin(ctx iris.Context) {
	response := service.UserLogin(ctx)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func UserLogout(ctx iris.Context) {
	response := service.UserLogout(ctx)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}
