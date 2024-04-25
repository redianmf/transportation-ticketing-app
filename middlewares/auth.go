package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redianmf/transportation-ticketing-app/database"
	"github.com/redianmf/transportation-ticketing-app/domain"
	"github.com/redianmf/transportation-ticketing-app/repository"
)

func ValidateAuth(c *gin.Context) {
	// Get JWT from header
	tokenString, err := getJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Decode JWT
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Check expire
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

			// Get user
			var user = domain.User{}
			existingUser, err := repository.GetUserById(database.DbConnection, user)
			if err != nil {
				panic(err)
			}

			if existingUser.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Set("userId", int(claims["sub"].(float64)))
	c.Next()
}

func getJWT(authString string) (token string, err error) {
	if authString == "" {
		return "", errors.New("please authenticate")
	}

	tokenArr := strings.Split(authString, " ")
	if len(tokenArr) != 2 {
		return "", errors.New("invalid credentials")
	}

	return tokenArr[1], nil
}
