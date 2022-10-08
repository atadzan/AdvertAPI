package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const(
	salt = "abcdefghlmn123456"
	tokenTLL = 12 * time.Hour
	signingKey = "qwertyuiop654321"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}

func(s *AuthService) CreateUser(user AdvertAPI.SignUpInput)(int, error){
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func(s *AuthService) GenerateToken(username, password string)(string, error){
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTLL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}


func generatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}