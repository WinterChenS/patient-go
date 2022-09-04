package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"winterchen.com/patient-go/patient/dao"
	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/helpers"
	"winterchen.com/patient-go/patient/models"
	"winterchen.com/patient-go/patient/requests"
	"winterchen.com/patient-go/patient/responses"
)

// add patient
func AddPatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		helpers.HandleValidatorError(c, err)
		return
	}
	if err := dao.AddPatient(&patient); err != nil {
		responses.Error(c, 400, 400, err.Error(), nil)
		return
	}
	responses.Success(c, 200, "success", nil)
}

// page patient
func PagePatient(c *gin.Context) {
	var pageRequest requests.PageRequest
	if err := c.ShouldBindQuery(&pageRequest); err != nil {
		helpers.HandleValidatorError(c, err)
		return
	}
	page, err := dao.PageGetPatients(pageRequest)
	if err != nil {
		responses.Error(c, 400, 400, err.Error(), nil)
		return
	}
	responses.Success(c, 200, "success", page)
}

// upload image to minio
func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		responses.Error(c, 500, 500, "upload file error", "system error")
		return
	}
	// generate uuid for file name
	fileName := uuid.New().String() + file.Filename
	// upload file to minio bucket
	ok := helpers.UploadFile(global.Configs.Minio.BucketName, fileName, fileObj, file.Size)
	if !ok {
		responses.Error(c, 401, 401, "upload file error", "")
		return
	}
	headerUrl := helpers.GetFileUrl(global.Configs.Minio.BucketName, fileName)
	if headerUrl == "" {
		responses.Success(c, 400, "upload file error", "get file url error")
		return
	}
	responses.Success(c, 200, "success", map[string]interface{}{
		"url":      headerUrl,
		"fileName": fileName,
	})
}
