package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) { 
	c.JSON(http.StatusOK, gin.H{"data": "token oke wkkw"})
}