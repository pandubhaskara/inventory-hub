package controller

import (
	"errors"
	"fmt"
	"net/http"

	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"
	"credential-auth/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginController struct{}

func (ctlr *LoginController) Index(c *gin.Context) {
	type rule struct {
		ClientID     int    `json:"client_id" form:"client_id" binding:"required"`
		ClientSecret string `json:"client_secret" form:"client_secret" binding:"required"`
		GrantType    string `json:"grant_type" form:"grant_type" binding:"required"`
	}

	var r rule

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	var client model.Client

	if result := database.Db.Where("id = ? AND secret=?", r.ClientID, r.ClientSecret).First(&client); result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Invalid client_id/client_secret",
		})
		c.Abort()
		return
	}

	if r.GrantType == "password" {
		grantTypePassword(c)
	} else if r.GrantType == "refresh_token" {
		grantTypeRefreshToken(c)
	} else {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		c.Abort()
		return
	}
}

func grantTypePassword(c *gin.Context) {
	type rule struct {
		ClientID int    `json:"client_id" form:"client_id" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	var r rule

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	var user model.User
	result := database.Db.Where("email = ?", r.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(401, gin.H{
				"message": "Wrong email/password",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(401, gin.H{
				"message": result.Error,
			})
			c.AbortWithStatus(500)
			return
		}
	}

	//comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"type":    "authentication",
			"message": "Wrong email/password",
		})
		c.Abort()
		return
	}

	//check user verification status
	if user.VerifiedAt == nil {
		c.JSON(401, gin.H{
			"type":    "user_not_verified",
			"message": "Please verify your account first.",
		})
		c.Abort()
		return
	}
	//create token
	token, err := service.GenerateToken(user, r.ClientID)
	if err != nil {
		helper.Logger.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, token)
}

func grantTypeRefreshToken(c *gin.Context) {
	type rule struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
	}

	var r rule

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	token, err := jwt.Parse(r.RefreshToken, func(token *jwt.Token) (interface{}, error) {
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

	database.Db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		var accessToken model.AccessToken
		var refreshToken model.RefreshToken

		result := tx.Where("id=? AND is_active=TRUE AND expired_at >= NOW()", claims["jti"]).First(&refreshToken)
		if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(401, gin.H{
				"message": "Refresh token failed",
				"type":    "refresh_token_failed",
			})
			c.Abort()
			return result.Error
		} else if result.Error != nil {
			helper.Logger.Error(result.Error)
			c.AbortWithStatus(500)
			return result.Error
		}

		refreshToken.IsActive = false
		tx.Save(&refreshToken)

		tx.Where("id = ?", claims["sub"]).First(&accessToken)
		accessToken.IsActive = false
		tx.Save(&accessToken)

		tx.Where("id=? ", accessToken.UserID).First(&user)

		newToken, err := service.GenerateToken(user, accessToken.ClientID)
		if err != nil {
			helper.Logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return result.Error
		}
		c.JSON(200, newToken)
		return nil
	})
}
