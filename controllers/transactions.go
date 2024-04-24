package controllers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redianmf/transportation-ticketing-app/database"
	"github.com/redianmf/transportation-ticketing-app/domain"
	"github.com/redianmf/transportation-ticketing-app/helpers"
	"github.com/redianmf/transportation-ticketing-app/repository"
)

func TransactionGate(c *gin.Context) {
	var user domain.User

	// Get user id
	userId := c.MustGet("userId").(int)
	user.ID = userId

	transaction, err := repository.GetLastTransactionByUserId(database.DbConnection, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	isNoTransaction := transaction == (domain.Transaction{})

	if isNoTransaction || transaction.Status != "GATE_IN" {
		GateIn(c)
	} else {
		GateOut(c, transaction)
	}
}

func GateIn(c *gin.Context) {
	var transaction domain.Transaction

	userId := c.MustGet("userId").(int)

	transModeId, err := strconv.Atoi(c.Param("transportationModeId"))
	if err != nil {
		panic(err)
	}

	pointId, err := strconv.Atoi(c.Param("pointId"))
	if err != nil {
		panic(err)
	}

	dateString := time.Now().Format("060102")
	randomNumberString := strconv.Itoa(rand.Int())

	transaction.UserId = userId
	transaction.TransportationModeId = transModeId
	transaction.TransactionReference = "TRX-" + dateString + randomNumberString
	transaction.Type = "DEBIT"
	transaction.Status = "GATE_IN"
	transaction.Amount = 0
	transaction.StartPoint = pointId
	transaction.EndPoint = pointId
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err = repository.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Passed GATE IN",
	})
}

func GateOut(c *gin.Context, lastTransaction domain.Transaction) {

	// Get params
	transModeId, err := strconv.Atoi(c.Param("transportationModeId"))
	if err != nil {
		panic(err)
	}

	pointId, err := strconv.Atoi(c.Param("pointId"))
	if err != nil {
		panic(err)
	}

	// Get Transportation mode used by user
	var transportationMode = domain.TransportationModes{
		ID: lastTransaction.TransportationModeId,
	}

	transportationData, err := repository.GetTransportationModeById(database.DbConnection, transportationMode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Transportation mode not found",
		})
		return
	}

	// Validation
	if transModeId != lastTransaction.TransportationModeId {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot gate out with different transportation mode",
		})
		return
	}

	// Update transaction
	var transaction = domain.Transaction{
		ID:        lastTransaction.ID,
		Status:    "GATE_OUT",
		Amount:    helpers.CalculateFare(transportationData, lastTransaction.StartPoint, pointId),
		EndPoint:  pointId,
		UpdatedAt: time.Now(),
	}

	err = repository.UpdateTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}

	// Deduct wallet
	var wallet = domain.Wallet{
		UserId: lastTransaction.UserId,
	}

	walletData, err := repository.GetWalletByUserId(database.DbConnection, wallet)
	if err != nil {
		panic(err)
	}

	wallet.Amount = walletData.Amount - transaction.Amount
	wallet.UpdatedAt = time.Now()

	err = repository.UpdateWalletByUserId(database.DbConnection, wallet)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Passed GATE OUT",
	})
}
