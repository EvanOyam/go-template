package common

import (
	"go-server-template/pkg/database"
	"go-server-template/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Result struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

var ErrorMap = map[int32]string{
	10001: "system error",
	10002: "invalid parameter",
	10003: "invalid user id",
}

func JsonError(code int, message string) *Result {
	return &Result{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
	}
}

func JsonSuccess(data interface{}) *Result {
	return &Result{
		ErrorCode: 0,
		Message:   "success",
		Data:      data,
	}
}

func HandleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logger.Infof("got signal [%s], exiting now", s)
		if err := server.Close(); nil != err {
			logger.Errorf("server close failed: " + err.Error())
		}

		database.CloseDB()

		logger.Infof("Exited")
		os.Exit(0)
	}()
}
