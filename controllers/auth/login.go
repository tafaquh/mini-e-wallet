package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) { 
	c.JSON(http.StatusOK, gin.H{"data": "login oke wkkw"})
}