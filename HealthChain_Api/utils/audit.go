package utils

import (
	"HealthChain_API/config"
	"HealthChain_API/models"
	"time"
)

func LogAudit(userID uint64, action string, details string) error {
	audit := models.Audit{
		UserID:    userID,
		Action:    action,
		Timestamp: time.Now(),
	}
	result := config.DB.Create(&audit)
	return result.Error
}
