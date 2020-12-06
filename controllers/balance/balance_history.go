package balance

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetBalanceHistory(c *gin.Context) { 
	c.JSON(http.StatusOK, gin.H{"data": "balance history oke wkkw"})
}