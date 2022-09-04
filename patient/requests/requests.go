package requests

import (
	"time"
)

type PageRequest struct {
	PageNum  int `form:"page_num" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

type PatientRequest struct {
	Name            string    `json:"name"`
	Brithday        time.Time `json:"brithday"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	Photo           string    `json:"photo"`
	DriverLicense   string    `json:"driver_license"`
	AppointmentTime time.Time `json:"appointment_time"`
}
