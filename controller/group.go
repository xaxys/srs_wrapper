package controller

import (
	"srs_wrapper/dao"
	"srs_wrapper/model"
	"srs_wrapper/util"

	"github.com/kataras/iris/v12"
)

func GetGroup(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	group := dao.GetGroupByID(id)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(model.ApiResponse(true, group, "操作成功"))
}

func CreateGroup(ctx iris.Context) {
	aul := &model.GroupJson{}
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
		u := dao.CreateGroup(aul, []*model.Permission{})
		ctx.StatusCode(iris.StatusOK)
		if u.ID == 0 {
			ctx.JSON(model.ApiResponse(false, u, "操作失败"))
		} else {
			ctx.JSON(model.ApiResponse(true, u, "操作成功"))
		}
	}
}

func UpdateGroup(ctx iris.Context) {
	aul := &model.GroupJson{}
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
		u := dao.UpdateGroup(aul, id)
		ctx.StatusCode(iris.StatusOK)
		if u.ID == 0 {
			ctx.JSON(model.ApiResponse(false, u, "操作失败"))
		} else {
			ctx.JSON(model.ApiResponse(true, u, "操作成功"))
		}
	}

}

func DeleteGroup(ctx iris.Context) {
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
	} else {
		dao.DeleteGroupByID(id)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ApiResponse(true, nil, "删除成功"))
	}
}

func GetAllGroups(ctx iris.Context) {
	aul := &model.AllGroupReq{}
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
		groups := dao.GetAllGroupsWithParam(aul.Name, aul.OrderBy, aul.Offset, aul.Limit)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(model.ApiResponse(true, groups, "操作成功"))
	}
}
