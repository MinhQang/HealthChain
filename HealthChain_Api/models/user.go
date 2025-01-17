package models

import "HealthChain_API/config"

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Role     string `json:"role"` // "doctor" or "patient"
}

func CreateUser(user *User) error {
	return config.DB.Create(user).Error
}

func GetUserByUsername(username string, user *User) error {
	return config.DB.Where("username = ?", username).First(user).Error
}

func GetUserByEmail(email string, user *User) error {
	return config.DB.Where("email = ?", email).First(user).Error
}
