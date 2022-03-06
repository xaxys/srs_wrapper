package controller

import (
	"fmt"
	"srs_wrapper/config"
	"srs_wrapper/dao"
	"srs_wrapper/model"

	"github.com/kataras/iris/v12"
)

func PostConnect(ctx iris.Context) {
	body := ctx.Values().Get("body").(*model.SrsCallbackReq)

	if body.Action != "on_connect" {
		ctx.Write(model.PermRejectedRes)
		return
	}

	if userID, ok := ctx.Values().Get("user_id").(uint); ok {
		dao.CreateClientWithUserID(body.ClientID, userID)
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %s] logined as [User %d]\n", body.ClientID, userID)
	} else if config.AppConfig.GetBool("app.guest") {
		dao.CreateGuestClient(body.ClientID)
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %s] logined as [Guest]\n", body.ClientID)
	} else {
		ctx.Write(model.PermRejectedRes)
	}
}

func PostClose(ctx iris.Context) {
	body := ctx.Values().Get("body").(*model.SrsCallbackReq)

	if body.Action != "on_close" {
		ctx.Write(model.PermRejectedRes)
		return
	}

	dao.DeleteClient(body.ClientID)
	ctx.Write(model.PermGrantedRes)
	fmt.Printf("[Client %s] logout\n", body.ClientID)
}

func PostPublish(ctx iris.Context) {
	body := ctx.Values().Get("body").(*model.SrsCallbackReq)

	if body.Action != "on_publish" {
		ctx.Write(model.PermRejectedRes)
		return
	}

	group := dao.GetGroupByClient(body.ClientID)
	if group == nil && config.AppConfig.GetBool("app.guest") {
		group = dao.GetGroupByID(dao.GuestGroupID)
	} else {
		ctx.Write(model.PermRejectedRes)
		fmt.Printf("[Client %s] as guest is forbidden to publish stream\n", body.ClientID)
	}
	if !dao.HasPermission(group, "callback.publish") {
		ctx.Write(model.PermRejectedRes)
		fmt.Printf("[Client %s] has no permission to publish stream\n", body.ClientID)
	} else {
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %s] start to publish stream\n", body.ClientID)
	}
}

func PostUnpublish(ctx iris.Context) {
	body := ctx.Values().Get("body").(*model.SrsCallbackReq)

	if body.Action != "on_unpublish" {
		ctx.Write(model.PermRejectedRes)
		return
	}

	ctx.Write(model.PermGrantedRes)
	fmt.Printf("[Client %s] stop to publish stream\n", body.ClientID)
}

func PostPlay(ctx iris.Context) {
	body := ctx.Values().Get("body").(*model.SrsCallbackReq)

	if body.Action != "on_play" {
		ctx.Write(model.PermRejectedRes)
		return
	}

	group := dao.GetGroupByClient(body.ClientID)
	if group == nil && config.AppConfig.GetBool("app.guest") {
		group = dao.GetGroupByID(dao.GuestGroupID)
	} else {
		ctx.Write(model.PermRejectedRes)
		fmt.Printf("[Client %s] as guest is forbidden to play stream\n", body.ClientID)
	}
	if !dao.HasPermission(group, "callback.play") {
		ctx.Write(model.PermRejectedRes)
		fmt.Printf("[Client %s] has no permission to play stream\n", body.ClientID)
	} else {
		ctx.Write(model.PermGrantedRes)
		fmt.Printf("[Client %s] start to play stream\n", body.ClientID)
	}
}

func PostStop(ctx iris.Context) {
	body := ctx.Values().Get("body").(*model.SrsCallbackReq)

	if body.Action != "on_stop" {
		ctx.Write(model.PermRejectedRes)
		return
	}

	ctx.Write(model.PermGrantedRes)
	fmt.Printf("[Client %s] stop to play stream\n", body.ClientID)
}
