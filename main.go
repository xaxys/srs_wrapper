package main

import (
	"github.com/kataras/iris/v12"

	"srs_wrapper/config"
	"srs_wrapper/initialize"
	"srs_wrapper/route"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel(config.AppConfig.GetString("app.loglevel"))
	initialize.InitDefaultData()
	route.Route(app)
	app.Listen(config.AppConfig.GetString("app.listen"))
}
