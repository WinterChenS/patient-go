/*
 * @Date: 2022-09-03 12:47:13
 * @Author: winterchen
 * @Description: TODO
 * @Version: 1.0
 * @LastEditTime: 2022-09-05 13:40:29
 */
package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/initialize"
	"winterchen.com/patient-go/patient/middlewares"
)

func main() {

	// init config
	initialize.InitConfig()

	// init routers
	Router := initialize.Routers()

	// init logger
	initialize.InitLogger()

	// init db
	initialize.InitDB()

	// init minio
	initialize.InitMinIO()

	// init snowflake
	if err := initialize.InitSnowflake(time.Now(), 1); err != nil {
		panic(err)
	}

	if err := initialize.InitTrans("en"); err != nil {
		panic(err)
	}

	color.Cyan("patient-go is running")

	// set logger middleware
	Router.Use(middlewares.LoggerForGin())

	// use cors
	Router.Use(middlewares.Cors())

	// start server
	err := Router.Run(fmt.Sprintf(":%d", global.Configs.Port))
	if err != nil {
		global.Log.Info("the patient-go ", zap.String("error", "run failed"))
	}

}
