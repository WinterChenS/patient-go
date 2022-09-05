package initialize

import (
	"fmt"

	"go.uber.org/zap"
	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/helpers"
)

func InitLogger() {
	// init zap config
	cfg := zap.NewDevelopmentConfig()
	// set log path
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Configs.LogsPath, helpers.GetNowFormatTodayTime()), //
		"stdout",
	}
	// create logger
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	// set global logger
	global.Log = logger
}
