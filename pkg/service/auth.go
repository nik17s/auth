package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const jwtTTL = 12 * time.Hour

type JWTClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func (s *Service) CreateUser(login, password, email, phone string) (int, error) {
	hashedPassword := getHashedPassword(password)

	return s.repository.CreateUser(login, hashedPassword, email, phone)
}

func (s *Service) GenerateJWT(login, password string) (string, error) {
	user, err := s.repository.GetUser(login, getHashedPassword(password))
	if err != nil {
		return "", err
	}

	token, err := generateJWT(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) ParseJWT(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New("invalid jwt")
	}
}

func getHashedPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func generateJWT(userId int) (string, error) {
	JWTKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
