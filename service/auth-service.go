package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Auth(c *gin.Context)
}

type jwtAuthService struct {
	secret string
}

type fakeAuthService struct {
}

func JwtAuthService() AuthService {
	secret := os.Getenv("AUTH_SECRET")
	return &jwtAuthService{secret}
}

func FakeAuthService() AuthService {
	return &fakeAuthService{}
}

func (j *fakeAuthService) Auth(c *gin.Context) {
	c.Next()
}

func (j *jwtAuthService) Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Athorization header is missing"})
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header"})
		return
	}
	token, err := jwt.Parse(authParts[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}
	c.Next()
}
