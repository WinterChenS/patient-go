package controllers

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/patient-go/patient/logic"
)

// add patient
func AddPatient(c *gin.Context) {
	logic.AddPatient(c)
}

// page patient
func PagePatient(c *gin.Context) {
	logic.PagePatient(c)
}

// upload image to minio
func UploadImage(c *gin.Context) {
	logic.UploadImage(c)
}
