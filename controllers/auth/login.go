package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tafaquh/mini-e-wallet/dto"
	"github.com/tafaquh/mini-e-wallet/database"
	"github.com/tafaquh/mini-e-wallet/models/user"
	"github.com/tafaquh/mini-e-wallet/services/jwt"
)

func Login(c *gin.Context) { 
	var credential dto.LoginCredentials
	err := c.ShouldBind(&credential)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Your credentials are invalid!"})
		return 
	}

	var token string
	isAuthenticated := Authenticate(credential.Email, credential.Password)
	if isAuthenticated {
		token = jwt.GenerateToken(credential.Email, true)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Your credentials are invalid!"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success login!", 
		"data": gin.H{
			"email":  credential.Email,
			"token": token,
		},
	})
}

func Authenticate(email string, password string) bool {
	db := database.DBConn

	var user_data user.User
	if res := db.Where("email = ?", email).Find(&user_data); res.RowsAffected == 0{
		return false
	}

	if err := user.VerifyPassword(user_data.Password, password); err != nil {
		return false
	}
	
	return true
}