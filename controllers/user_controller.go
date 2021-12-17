package controllers

import (
	"go-server-template/pkg/common"
	"go-server-template/services"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

type User struct {
	id int
}

func (c *UserController) Get() *common.Result {
	id, err := c.Ctx.URLParamInt64("id")
	if err != nil {
		return common.JsonError(10002, common.ErrorMap[10002])
	}

	user := services.UserService.GetUserInfo(id)

	if user == nil || !user.Username.Valid {
		return common.JsonError(10003, common.ErrorMap[10003])
	}

	return common.JsonSuccess(iris.Map{
		"username":    user.Username.String,
		"nickname":    user.Nickname,
		"description": user.Description,
	})
}
