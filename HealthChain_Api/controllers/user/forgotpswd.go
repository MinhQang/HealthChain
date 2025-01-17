package user

import (
	"HealthChain_API/config"
	"HealthChain_API/models"
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

var otpMap = map[string]string{}

func ForgotPwController(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Yêu cầu không hợp lệ", http.StatusBadRequest)
		return
	}

	var user models.User
	config.DB.Where("email = ?", request.Email).First(&user)
	if user.ID == 0 {
		http.Error(w, "Không tìm thấy email", http.StatusNotFound)
		return
	}

	otp := generateOTP()
	otpMap[user.Email] = otp

	err = sendEmail(user.Email, otp)
	if err != nil {
		http.Error(w, "Lỗi khi gửi email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"otp": "đã được gửi tới email"})
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(1000000)
	return fmt.Sprintln("%06d", otp)
}

func sendEmail(email string, otp string) error {
	from := "your-email@example.com"
	password := "your-email-password"

	msg := "Subject: Reset Password OTP\n\nYour OTP is: " + otp
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(msg))
	return err
}

func ResetPwController(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
		NewPW string `json:"newPW"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Yêu cầu không hợp lệ", http.StatusBadRequest)
		return
	}
	storedOTP, exists := otpMap[request.Email]
	if !exists || storedOTP != request.OTP {
		http.Error(w, "OTP không hợp lệ", http.StatusUnauthorized)
		return
	}

	var user models.User
	config.DB.Where("email = ?", request.Email).First(&user)
	if user.ID == 0 {
		http.Error(w, "Không tìm thấy Email", http.StatusNotFound)
		return
	}

	user.Password = hashPassword(request.NewPW)
	config.DB.Save(&user)

	delete(otpMap, request.Email) // Xóa OTP sau khi sử dụng

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successful"})
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
