package middlewares

import (
	"go-server-template/pkg/logger"
	"time"

	"github.com/kataras/iris/v12"
)

func Log(ctx iris.Context) {
	referrer := ctx.GetReferrer()
	url := ctx.Request().URL
	ua := ctx.Request().UserAgent()
	logger.Infof("Get the requese at [%v], url is: [%v], referrer is: [%v], ua is: [%v]\n", time.Now(), url, referrer, ua)
	ctx.Next()
}
