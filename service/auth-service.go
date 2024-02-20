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
}

type jwtAuthService struct {
	secret string
	apiKey string
}

type fakeAuthService struct {
}

func JwtAuthService() AuthService {
	secret := os.Getenv("AUTH_SECRET")
	apiKey := os.Getenv("API_KEY")
	return &jwtAuthService{secret, apiKey}
}

func FakeAuthService() AuthService {
	return &fakeAuthService{}
}

func (j *fakeAuthService) HeaderAuth(c *gin.Context) {
	c.Next()
}

func (j *jwtAuthService) tokenAuth(c *gin.Context, authHeader string) {
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization header"})
		return
	}

	token, err := j.getToken(authParts[1])
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	c.Next()
}

func (j *jwtAuthService) HeaderAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		j.tokenAuth(c, authHeader)
		return
	}
	authHeader = c.GetHeader("ApiKey")
	if authHeader != "" {
		if j.apiKey != authHeader {
			c.AbortWithStatusJSON(401, gin.H{"error": "wrong api key"})
			return
		}
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, gin.H{"error": "Athorization header is missing"})

}

func (j *jwtAuthService) getToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
