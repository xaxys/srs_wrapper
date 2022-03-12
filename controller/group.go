package controller

import (
	"srs_wrapper/model"
	"srs_wrapper/service"

	"github.com/kataras/iris/v12"
)

func GetGroup(ctx iris.Context) {
	id := ctx.Values().Get("user_id").(uint)
	response := service.GetGroup(id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func CreateGroup(ctx iris.Context) {
	aul := &model.GroupJson{}
	if err := ctx.ReadJSON(&aul); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	id := ctx.Values().Get("user_id").(uint)
	response := service.CreateGroup(id, aul)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func UpdateGroup(ctx iris.Context) {
	aul := &model.GroupJson{}
	if err := ctx.ReadJSON(&aul); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	qid, _ := ctx.Params().GetUint("id")
	id := ctx.Values().Get("user_id").(uint)
	response := service.UpdateGroup(qid, id, aul)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func DeleteGroup(ctx iris.Context) {
	qid, _ := ctx.Params().GetUint("id")
	id := ctx.Values().Get("user_id").(uint)
	response := service.DeleteGroup(qid, id)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}

func GetAllGroups(ctx iris.Context) {
	req := &model.AllGroupReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		response := model.ErrorInvalidData(err)
		ctx.StatusCode(response.Code)
		ctx.JSON(response)
		return
	}
	response := service.GetAllGroups(req)
	ctx.StatusCode(response.Code)
	ctx.JSON(response)
}
