package log

import "go.uber.org/zap"

var L *zap.Logger

func Init(env string) {
	var err error
	if env == "prod" {
		L, err = zap.NewProduction()
	} else {
		L, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
}
