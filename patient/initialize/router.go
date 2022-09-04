package initialize

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/patient-go/patient/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("/api/v1/")
	router.PatientRouter(ApiGroup)
	return r
}
