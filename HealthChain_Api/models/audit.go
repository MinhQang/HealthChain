package models

import (
	"HealthChain_API/config"
	"time"
)

type Audit struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Action    string    `json:"action" gorm:"not null"`
	UserID    uint64    `json:"userId" gorm:"not null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
}

func CreateAudit(audit *Audit) error {
	audit.Timestamp = time.Now()
	return config.DB.Create(audit).Error
}
