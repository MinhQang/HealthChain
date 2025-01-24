package utils

import (
	"math/rand"
	"net/smtp"
	"time"
)

var OtpMap = map[string]string{}

func GenerateOTP() string {
	const charset = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(otp)
}

func SendEmail(otp string, email string) error {
	from := "your-email@example.com"
	password := "your-email-password"

	// Email message
	msg := "Subject: Reset Password OTP\n\nYour OTP is: " + otp
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(msg))
	return err
}
