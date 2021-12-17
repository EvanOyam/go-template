package main

import (
	"flag"
	"go-server-template/controllers"
	"go-server-template/models"
	"go-server-template/pkg/common"
	"go-server-template/pkg/config"
	"go-server-template/pkg/database"
	"go-server-template/pkg/logger"
	"net/http"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	irisLogger "github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

var env = flag.String("env", "dev", "运行模式")

func init() {
	flag.Parse()
	var configPath string
	if *env == "prod" {
		configPath = "./config/prod.yaml"
	} else {
		configPath = "./config/dev.yaml"
	}

	// 初始化配置
	if err := config.InitConfig(configPath); err != nil {
		logger.Error(err)
		panic("Initial configuration failed!")
	}

	logPath := config.Get("LogFile")
	mySqlUrl := config.Get("MySqlUrl")

	// 初始化日志
	if err := logger.InitLog(logPath.(string)); err != nil {
		logger.Error(err)
		panic("Initial logger failed!")
	}

	// 初始化数据库
	if err := database.OpenDB(mySqlUrl.(string), logPath.(string), 10, 20, models.Models...); err != nil {
		logger.Error(err)
		panic("Initial database failed!")
	}

	logger.Info("Service initialization successful.")
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.Use(irisLogger.New())

	// 跨域处理
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	app.AllowMethods(iris.MethodOptions)

	// 异常 http 状态码处理
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.JSON(common.JsonError(10001, common.ErrorMap[10001]))
	})

	// 404 兜底
	app.Any("/", func(i iris.Context) {
		_, _ = i.HTML("<h1>Powered by evan</h1>")
	})

	// 初始化 controllers
	controllers.Init(app)

	server := &http.Server{Addr: ":" + config.Get("Port").(string)}

	// 处理进程异常退出
	common.HandleSignal(server)
	err := app.Run(iris.Server(server), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))

	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
}
