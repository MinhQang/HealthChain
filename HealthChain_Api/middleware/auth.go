package middleware

import (
	"HealthChain_API/utils"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

type EmailRequest struct {
	Email string `json:"email"`
}

var otpStore = make(map[string]string)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var emailRequest EmailRequest
		err = json.NewDecoder(r.Body).Decode(&emailRequest)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusUnauthorized)
			return
		}

		otp := utils.GenerateOTP()
		otpStore[emailRequest.Email] = otp
		go func() {
			err := utils.SendEmail(emailRequest.Email, otp)
			if err != nil {
				log.Println("Error sending email:", emailRequest.Email, err)
			} else {
				err := utils.LogAudit(claims.UserID, "SendOTP", "Sent OTP to "+emailRequest.Email)
				if err != nil {
					log.Println("Failed to log audit: ", err)
				}
			}
		}()

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func VerifyOTP(email, otp string) bool {
	if storedOTP, ok := otpStore[email]; ok {
		if storedOTP == otp {
			delete(otpStore, email)
			return true
		}
	}
	return false
}
