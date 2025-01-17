package migration

import (
	"HealthChain_API/config"
	"HealthChain_API/models"
	"log"
)

func Migration() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Patient{}, &models.Audit{}, &models.Access{}, &models.EmergencyContact{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
