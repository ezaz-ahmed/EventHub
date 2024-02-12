package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	secretKey []byte
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

type Claims struct {
	Email  string `json:"email"`
	UserID int64  `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(email string, userId int64) (string, error) {
	claims := &Claims{
		Email:  email,
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userId := claims.UserID

	return userId, nil
}
