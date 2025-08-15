package log

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init(env string) {
	var err error
	if env == "production" || env == "prod" {
		Logger, err = zap.NewProduction()
	} else {
		Logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("cannot init logger: " + err.Error())
	}
}

func Sync() {
	_ = Logger.Sync()
}
