package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"online_lib_api"
	"online_lib_api/internal/storage"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "sfuahfuisherw8234r234"
	signingKey = "vghfdijkgdrjkh4trherfbnhjkdfs345r"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{storage: storage}
}

func (s *AuthService) CreateUser(user online_lib_api.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.storage.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})
		

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}