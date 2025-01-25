package utils

import (
	"HealthChain_API/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte("secret-key")

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

//func ParseJWT(tokenString string) (*Claims, error) {
//	claims := &Claims{}
//	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//		return JwtKey, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if !token.Valid {
//		return nil, jwt.ErrSignatureInvalid
//	}
//	return claims, nil
//}
