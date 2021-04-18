package router

import (
	"gis-hcsb/api"
	"gis-hcsb/middleware"
	"github.com/kataras/iris/v12"
)

func App() *iris.Application {
	std := iris.Default()
	std.Logger().SetLevel("debug")
	std.Use(middleware.Cors())

	std.Post("/register", api.Register)
	std.Post("/login", api.Login)
	std.Post("/history", api.History)
	std.Post("/upload", api.Upload)
	std.Post("/gpass", api.GpAss)

	return std
}
