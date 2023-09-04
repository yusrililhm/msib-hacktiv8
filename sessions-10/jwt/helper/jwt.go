package helper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseTkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := parseTkn.SignedString([]byte("secret-key"))

	return token
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	var err error = errors.New("signin to procced")

	headerTkn := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerTkn, "Bearer")

	if !bearer {
		return nil, err
	}

	stringTkn := strings.Split(headerTkn, " ")[1]

	token, _ := jwt.Parse(stringTkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("secret-key"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
