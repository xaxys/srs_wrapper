package controller

import (
	"fmt"
	"srs_wrapper/config"
	"srs_wrapper/dao"
	"srs_wrapper/model"

	"github.com/kataras/iris/v12"
)

func PostConnect(ctx iris.Context) {
	req := model.OnConnectReq{}
	ctx.ReadJSON(&req)

	if req.Action != "on_connect" {
		ctx.Write(model.PermRejectedRes)
	}

	if userID, ok := ctx.Values().Get("user_id").(uint); ok {
		dao.CreateClientWithUserID(req.ClientID, userID)
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %v] logined as [User %v]", req.ClientID, userID)
	} else if config.AppConfig.GetBool("app.guest") {
		dao.CreateGuestClient(req.ClientID)
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %v] logined as [Guest]", req.ClientID)
	} else {
		ctx.Write(model.PermRejectedRes)
	}
}

func PostClose(ctx iris.Context) {
	req := model.OnCloseReq{}
	ctx.ReadJSON(&req)

	if req.Action != "on_close" {
		ctx.Write(model.PermRejectedRes)
	}

	dao.DeleteClient(req.ClientID)
	ctx.Write(model.PermGrantedRes)
	fmt.Printf("[Client %v] logout", req.ClientID)
}

func PostPublish(ctx iris.Context) {
	req := model.OnPublishReq{}
	ctx.ReadJSON(&req)

	if req.Action != "on_publish" {
		ctx.Write(model.PermRejectedRes)
	}

	group := dao.GetGroupByClient(req.ClientID)
	if group != nil && !dao.HasPermission(group, "callback.publish") {
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %v] start to publish stream", req.ClientID)
	} else {
		ctx.Write(model.PermRejectedRes)
		fmt.Printf("[Client %v] has no permission to publish stream", req.ClientID)
	}
}

func PostUnpublish(ctx iris.Context) {
	req := model.OnUnpublishReq{}
	ctx.ReadJSON(&req)

	if req.Action != "on_unpublish" {
		ctx.Write(model.PermRejectedRes)
	}

	ctx.Write(model.PermGrantedRes)
	fmt.Printf("[Client %v] stop to publish stream", req.ClientID)
}

func PostPlay(ctx iris.Context) {
	req := model.OnPlayReq{}
	ctx.ReadJSON(&req)

	if req.Action != "on_play" {
		ctx.Write(model.PermRejectedRes)
	}

	group := dao.GetGroupByClient(req.ClientID)
	if group != nil && !dao.HasPermission(group, "callback.play") {
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %v] start to play stream", req.ClientID)
	} else {
		ctx.Write(model.PermRejectedRes)
		fmt.Printf("[Client %v] has no permission to play stream", req.ClientID)
	}
}

func PostStop(ctx iris.Context) {
	req := model.OnStopReq{}
	ctx.ReadJSON(&req)

	if req.Action != "on_stop" {
		ctx.Write(model.PermRejectedRes)
	}

	ctx.Write(model.PermGrantedRes)
	fmt.Printf("[Client %v] stop to play stream", req.ClientID)
}
