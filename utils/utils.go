package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

func CreateUUIDStr() (string, error) {
	// generate an uuid and return it as a string
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u4.String(), nil
}

// HTTP Response
func SendInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "failed",
		"message": err.Error(),
	})
	c.Abort()
}

func SendUnauthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "not authorized",
		"error":   err.Error(),
	})
	c.Abort()
}

func SendBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "failed",
		"message": err.Error(),
	})
	c.Abort()
}

func SendSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func SendData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
