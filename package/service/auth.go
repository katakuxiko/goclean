package service

import (
	"crypto/sha1"
	"fmt"
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/structure"
)
const (
	salt = "1easg34tgregv"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	toketTTL = 120 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func(s *AuthService) CreateUser(user structure.User)(int, error){
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) GenerateToken(username,password string)(string, error){
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(toketTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
func (s *AuthService)ParseToken(accessToken string)(int, error){
	token,err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token)(interface{},error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey),nil
	})
	if err != nil{
		return 0,err
	}

	claims,ok := token.Claims.(*tokenClaims)
	if !ok{
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
 
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}