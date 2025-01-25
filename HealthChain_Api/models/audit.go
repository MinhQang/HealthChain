package models

import (
	"time"
)

type Audit struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Action    string    `json:"action" gorm:"type:varchar(255)"`
	UserID    uint64    `json:"userId" gorm:"index"`
	Timestamp time.Time `json:"timestamp" gorm:"autoCreateTime"`
	Details   string    `gorm:"type:text"`
}

//func CreateAudit(audit *Audit) error {
//	audit.Timestamp = time.Now()
//	return config.DB.Create(audit).Error
//}
