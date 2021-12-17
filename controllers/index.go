package controllers

import (
	"go-server-template/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Init(app *iris.Application) {
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middlewares.Log)
		m.Party("/user").Handle(new(UserController))
	})
}
