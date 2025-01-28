package service

import (
	"MeetEnjoy"
	"MeetEnjoy/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

const TokenTTL = 12 * time.Hour

type AuthService struct {
	repos repository.Authorization
}

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (as *AuthService) CreateUser(user MeetEnjoy.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.repos.CreateUser(user)
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.repos.GetUser(username, generatePasswordHash(password))
	if err != nil {
		log.Print("getUser")
		return "", err
	}

	claims := CustomClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTTL)), // Время истечения токена
			IssuedAt:  jwt.NewNumericDate(time.Now()),               // Время выпуска токена
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		log.Print("errorSignedString")
		return "", err
	}

	return tokenString, nil
}

func (as *AuthService) ParseToken(tokenString string) (int, error) {
	signingKey := []byte(os.Getenv("SIGNING_KEY"))

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return signingKey, nil
	})

	if err != nil {
		return 0, err
	}

	if token.Valid {
		return claims.UserId, nil
	}

	return 0, fmt.Errorf("invalid token")
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
