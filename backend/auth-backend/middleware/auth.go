package middleware

import (
	"encoding/json"
	"net/http"

	"credential-auth/database"
	"credential-auth/model"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	if len(c.Request.Header["X-User"]) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var xUser map[string]interface{}

	json.Unmarshal([]byte(c.Request.Header["X-User"][0]), &xUser)

	var data model.User
	database.Db.Where("email = ?", xUser["email"]).First(&data)

	c.Set("account", data)
	c.Next()
}
