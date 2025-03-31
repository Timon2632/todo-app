package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Timon2632/todo-app"
	"github.com/Timon2632/todo-app/pkg/repositoty"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "ahrghagfesf1243124wqadwafd"
	singingKey = "wqfewgrreq12`41512%@!112ewtewqw"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthServise struct {
	repo repositoty.Authorzation
}

func NewAuthServise(repo repositoty.Authorzation) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthServise) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(singingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
