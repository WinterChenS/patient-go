package initialize

import (
	"winterchen.com/patient-go/patient/config"
	"winterchen.com/patient-go/patient/global"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// init viper
	v := viper.New()
	// set config file
	v.SetConfigFile("./config-dev.yml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.Config{}
	//int serverConfig
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// set global config
	global.Configs = serverConfig
	color.Blue("init config", global.Configs.LogsPath)
}
