package router

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/patient-go/patient/controllers"
)

// patient router
func PatientRouter(Router *gin.RouterGroup) {
	patient := Router.Group("/patient")
	{
		patient.GET("/page", controllers.PagePatient)
		patient.POST("", controllers.AddPatient)
		patient.POST("/file", controllers.UploadImage)
	}
}
