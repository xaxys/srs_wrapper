package route

import (
	"srs_wrapper/controller"
	"srs_wrapper/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

// Route ...
func Route(app *iris.Application) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.CORS)
	app.AllowMethods(iris.MethodOptions)

	app.PartyFunc("/", func(home iris.Party) {
		home.HandleDir("/", "./assets")

		home.Get("/", func(ctx iris.Context) {
			ctx.Redirect("/index.html")
		})

		home.PartyFunc("/callback", func(callback iris.Party) {
			callback.Use(middleware.TcUrlExtractor, middleware.TokenValidator)
			callback.Post("/connect", controller.PostConnect)
			callback.Post("/close", controller.PostClose)
			callback.Post("/publish", controller.PostPublish)
			callback.Post("/unpublish", controller.PostUnpublish)
			callback.Post("/play", controller.PostPlay)
			callback.Post("/stop", controller.PostStop)
		})

		home.Post("/login", controller.UserLogin)
		home.PartyFunc("/", func(account router.Party) {
			account.Use(middleware.HeaderExtractor, middleware.TokenValidator, middleware.Interceptor)
			account.Get("/logout", controller.UserLogout)

			account.PartyFunc("/user", func(user router.Party) {
				user.Get("/all", controller.GetAllUsers)
				user.Get("/{id:uint}", controller.GetUser)
				user.Post("/", controller.CreateUser)
				user.Put("/{id:uint}", controller.UpdateUser)
				user.Delete("/{id:uint}", controller.DeleteUser)
				user.Get("/", controller.GetProfile)
			})
			account.PartyFunc("/group", func(group router.Party) {
				group.Get("/all", controller.GetAllGroups)
				group.Get("/{id:uint}", controller.GetGroup)
				group.Post("/", controller.CreateGroup)
				group.Put("/{id:uint}", controller.UpdateGroup)
				group.Delete("/{id:uint}", controller.DeleteGroup)
			})
			account.PartyFunc("/permission", func(perm router.Party) {
				perm.Get("/all", controller.GetAllPermissions)
				perm.Get("/{id:uint}", controller.GetPermission)
				perm.Post("/", controller.CreatePermission)
				perm.Put("/{id:uint}", controller.UpdatePermission)
				perm.Delete("/{id:uint}", controller.DeletePermission)
			})
		})
	})
}
