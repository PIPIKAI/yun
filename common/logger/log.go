package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// Logger
var Logger *zap.SugaredLogger

// InitLogger
func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败 %v", err))
	}
	Logger = logger.Sugar()
}
