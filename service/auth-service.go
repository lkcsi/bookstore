package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	HeaderAuth(c *gin.Context)
	CookieAuth(c *gin.Context)
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

func (j *fakeAuthService) HeaderAuth(c *gin.Context) {
	c.Next()
}

func (j *fakeAuthService) CookieAuth(c *gin.Context) {
	c.Next()
}

func (j *jwtAuthService) CookieAuth(c *gin.Context) {

	auth, err := c.Cookie("auth")
	if err != nil {
		fmt.Println(err.Error())
	}
	token, err := j.getToken(auth)
	if err != nil {
		c.AbortWithStatus(403)
	}
	if err != nil || !token.Valid {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}

func (j *jwtAuthService) HeaderAuth(c *gin.Context) {
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
	token, err := j.getToken(authParts[1])

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}
	c.Next()
}

func (j *jwtAuthService) getToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
