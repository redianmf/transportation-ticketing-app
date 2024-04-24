package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redianmf/transportation-ticketing-app/database"
	"github.com/redianmf/transportation-ticketing-app/domain"
	"github.com/redianmf/transportation-ticketing-app/repository"
)

func GetAllTransportationMode(c *gin.Context) {
	trModes, err := repository.GetAllTransportationMode(database.DbConnection)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": trModes,
		})
	}

}

func GetTransportationModeById(c *gin.Context) {
	var trMode domain.TransportationModes

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	trMode.ID = id

	trModeRes, err := repository.GetTransportationModeById(database.DbConnection, trMode)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": trModeRes,
		})
	}
}

func InsertTransportationMode(c *gin.Context) {
	var trMode domain.TransportationModes

	err := c.ShouldBindJSON(&trMode)
	if err != nil {
		panic(err)
	}

	trMode.CreatedAt = time.Now()
	trMode.UpdatedAt = time.Now()

	err = repository.InsertTransportationMode(database.DbConnection, trMode)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new transportation mode",
	})
}

func UpdateTransportationMode(c *gin.Context) {
	var trMode domain.TransportationModes

	err := c.ShouldBindJSON(&trMode)
	if err != nil {
		panic(err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	trMode.ID = id
	trMode.UpdatedAt = time.Now()

	err = repository.UpdateTransportationMode(database.DbConnection, trMode)
	if err != nil {
		panic(err)

	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update transportation mode",
	})
}

func DeleteTransportationMode(c *gin.Context) {
	var trMode domain.TransportationModes

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	trMode.ID = id

	err = repository.DeleteTransportationMode(database.DbConnection, trMode)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Transportation Mode",
	})
}
