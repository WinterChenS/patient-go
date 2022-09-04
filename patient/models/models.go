package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/helpers"
	"winterchen.com/patient-go/patient/requests"
)

type BaseModel struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
}

type User struct {
	BaseModel
	Username string
	Password string
	Email    string
	Status   int
}

type Patient struct {
	BaseModel
	Name            string    `json:"name" binding:"required"`
	Brithday        time.Time `json:"brithday" binding:"required"`
	Phone           string    `json:"phone" binding:"required"`
	Address         string    `json:"address"`
	Photo           string    `json:"photo"`
	DriverLicense   string    `json:"driver_license"`
	AppointmentTime time.Time `json:"appointment_time" binding:"required"`
}

type Page struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

/*
	user operations
*/

// get user by name
func (user *User) GetUserByName(name string) error {
	return global.DB.Where("username = ?", name).First(&user).Error
}

// add user
func (user *User) AddUser() error {
	return global.DB.Create(&user).Error
}

/*
	patient operations
*/

// add patient
func (patient *Patient) AddPatient() error {
	patient.ID = helpers.GenID()
	patient.CreatedAt = time.Now()
	return global.DB.Create(&patient).Error
}

// get patient by id
func (patient *Patient) GetPatientById(id int) error {
	return global.DB.Where("id = ?", id).First(&patient).Error
}

// page query
func PageGetPatients(pageRequest requests.PageRequest) (page *Page, err error) {
	var patients []Patient
	if err = global.DB.Order("created_at desc").Offset((pageRequest.PageNum - 1) * pageRequest.PageSize).Limit(pageRequest.PageSize).Find(&patients).Error; err == nil {
		var total int
		page = &Page{}
		global.DB.Model(&Patient{}).Count(&total)
		page.Total = total
		page.List = patients
		return page, err
	}
	return nil, err
}
