package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redianmf/transportation-ticketing-app/database"
	"github.com/redianmf/transportation-ticketing-app/domain"
	"github.com/redianmf/transportation-ticketing-app/repository"
)

func GetWalletByUserId(c *gin.Context) {
	// Get user id
	userId := c.MustGet("userId").(int)

	var wallet domain.Wallet
	wallet.UserId = userId

	// Get wallet
	walletData, err := repository.GetWalletByUserId(database.DbConnection, wallet)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": walletData,
	})

}

func UpdateWalletByUserId(c *gin.Context) {
	// Get user id
	userId := c.MustGet("userId").(int)

	// Get request body
	var wallet domain.Wallet

	err := c.ShouldBindJSON((&wallet))
	if err != nil {
		panic(err)
	}

	if wallet.Amount <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please verify the top up amount",
		})

		return
	}

	wallet.UserId = userId
	wallet.UpdatedAt = time.Now()

	err = repository.UpdateWalletByUserId(database.DbConnection, wallet)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success top up wallet",
	})

}
