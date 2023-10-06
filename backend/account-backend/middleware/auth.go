package middleware

import (
	"account-backend/helper"
	"encoding/json"
	"errors"
	"net/http"

	"account-backend/database"
	"account-backend/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(c *gin.Context) {
	if len(c.Request.Header["X-User"]) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var xUser map[string]interface{}

	json.Unmarshal([]byte(c.Request.Header["X-User"][0]), &xUser)

	var data model.Account
	result := database.Db.Where("email = ?", xUser["email"]).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			helper.Logger.Error(result.Error.Error())
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Account not found.",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	c.Set("account", data)
	c.Next()
}
