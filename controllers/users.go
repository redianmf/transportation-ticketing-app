package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redianmf/transportation-ticketing-app/database"
	"github.com/redianmf/transportation-ticketing-app/domain"
	"github.com/redianmf/transportation-ticketing-app/repository"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Get request body
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Incomplete data",
		})
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = string(hashed)

	// Check if user exist
	existingUser, err := repository.GetUserByEmail(database.DbConnection, user)

	if existingUser != (domain.User{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User already registered",
		})
		return
	}

	if err != nil {
		panic(err)
	}

	// Create User
	userId, err := repository.InsertUser(database.DbConnection, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create user",
		})
		return
	}

	var wallet = domain.Wallet{
		UserId:    userId,
		Amount:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repository.InsertWallet(database.DbConnection, wallet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success create user",
	})
}

func Login(c *gin.Context) {
	// Get request body
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	if user.Email == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please enter username and password",
		})
		return
	}

	// Check credentials
	existingUser, err := repository.GetUserByEmail(database.DbConnection, user)
	if existingUser == (domain.User{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	if err != nil {
		panic(err)
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		panic(err)
	}

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"user_id":  existingUser.ID,
		"username": existingUser.Username,
		"email":    existingUser.Email,
		"token":    tokenString,
	})
}
