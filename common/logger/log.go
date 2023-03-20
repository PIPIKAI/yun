package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败 %v", err))
	}
	Logger = logger.Sugar()
}
