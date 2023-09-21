package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/model"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type IntrospectController struct{}

func (ctlr *IntrospectController) Index(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(config.Jwt.Signature), nil
	})
	if err != nil {
		helper.Logger.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		helper.Logger.Error("claims failed")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var accessToken model.AccessToken

	result := database.Db.Where("id=? AND is_active=TRUE AND expired_at >= NOW()", claims["jti"]).First(&accessToken)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(401, gin.H{
			"message": "Access token failed",
			"type":    "access_token_failed",
		})
		c.Abort()
		return
	} else if result.Error != nil {
		c.AbortWithStatus(500)
		return
	}

	var data model.User

	database.Db.Where("id = ?", claims["sub"]).First(&data)

	jsonString, _ := json.Marshal(data)
	c.Header("X-User", string(jsonString))
	c.JSON(200, nil)
}
