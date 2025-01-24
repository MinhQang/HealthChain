package controllers

import (
	"HealthChain_API/middleware"
	"HealthChain_API/models"
	"HealthChain_API/utils"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type OTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Đầu vào không hợp lệ", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := models.GetUserByUsername(credentials.Username, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Người dùng không tồn tại", http.StatusUnauthorized)
		} else {
			http.Error(w, "Lỗi khi lấy dữ liệu người dùng", http.StatusInternalServerError)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Sai mật khẩu", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Không thể tạo token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Không hợp lệ", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Lỗi mã hóa mật khẩu", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := models.CreateUser(&user); err != nil {
		http.Error(w, "Lỗi khi tạo tài khoản", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func VerifyOTPController(w http.ResponseWriter, r *http.Request) {
	var otpRequest OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&otpRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if middleware.VerifyOTP(otpRequest.Email, otpRequest.OTP) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OTP verified successfully"))
	} else {
		http.Error(w, "OTP verification failed", http.StatusUnauthorized)
	}
}
