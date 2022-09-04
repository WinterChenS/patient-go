package dao

import (
	"time"

	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/helpers"
	"winterchen.com/patient-go/patient/models"
	"winterchen.com/patient-go/patient/requests"
)

// add patient
func AddPatient(patient *models.Patient) error {
	patient.ID = helpers.GenID()
	patient.CreatedAt = time.Now()
	return global.DB.Create(&patient).Error
}

// get patient by id
func GetPatientById(id int) (patient *models.Patient, err error) {
	if err = global.DB.Where("id = ?", id).First(&patient).Error; err == nil {
		return patient, nil
	}
	return nil, err
}

// page query
func PageGetPatients(pageRequest requests.PageRequest) (page *models.Page, err error) {
	var patients []models.Patient
	if err = global.DB.Order("created_at desc").Offset((pageRequest.PageNum - 1) * pageRequest.PageSize).Limit(pageRequest.PageSize).Find(&patients).Error; err == nil {
		var total int
		page = &models.Page{}
		global.DB.Model(&models.Patient{}).Count(&total)
		page.Total = total
		page.List = patients
		return page, err
	}
	return nil, err
}
