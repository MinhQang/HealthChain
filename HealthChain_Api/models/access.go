package models

import (
	"HealthChain_API/config"
)

type Access struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	UserID    uint64 `json:"userId" gorm:"not null"`
	PatientID uint64 `json:"patientId" gorm:"not null"`
	GrantedBy uint64 `json:"grantedBy" gorm:"not null"`
}

func CreateAccess(access *Access) error {
	return config.DB.Create(access).Error
}

func GetAccessByID(userID uint64, access *[]Access) error {
	return config.DB.Where("user_id = ?", userID).Find(access).Error
}

func GetAccessByPatientID(patientID uint64, access *Access) error {
	return config.DB.Where("patient_id = ?", patientID).Find(access).Error
}

func DeleteAccess(access *Access) error {
	return config.DB.Delete(access).Error
}
