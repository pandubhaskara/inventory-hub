package controller

import (
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmailVerificationController struct{}

func (ctrl *EmailVerificationController) Index(c *gin.Context) {
	token := c.Param("token")

	var v model.VerificationToken

	result := database.Db.Where("token=? AND is_active=? AND expired_at>= NOW()", token, true).First(&v)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invalid token.",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error,
			})
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Valid token.",
	})
}

func (ctrl *EmailVerificationController) Store(c *gin.Context) {
	type rule struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	var r rule

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	var v model.VerificationToken

	result := database.Db.Where("token=? AND is_active=? AND expired_at>= NOW()", r.Token, true).First(&v)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invalid token.",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error,
			})
			c.Abort()
			return
		}
	}

	//Update user verified status
	update := database.Db.Where("id = ?", v.UserID).Model(&model.User{}).Updates(map[string]interface{}{"verified_at": time.Now()})
	if update.Error != nil {
		helper.Logger.Error(update.Error)
		c.JSON(500, gin.H{
			"type":    "email_verification",
			"message": update.Error,
		})
		c.Abort()
		return
	}

	//update verification token status
	t := time.Now()

	v.IsActive = false
	v.UsedAt = &t

	deactiveToken := database.Db.Save(&v)
	if deactiveToken.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invalid token.",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error,
			})
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully verify your account.",
	})
}

func (ctrl *EmailVerificationController) SetPassword(c *gin.Context) {
	type rule struct {
		Token           string `json:"token" form:"token" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	var r rule

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	if r.Password != r.ConfirmPassword {
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": "Password and confirm password not matched.",
		})
		c.Abort()
		return
	}

	var v model.VerificationToken

	result := database.Db.Where("token=? AND is_active=? AND expired_at>= NOW()", r.Token, true).First(&v)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invalid token.",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error,
			})
			c.Abort()
			return
		}
	}

	// Store new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.Logger.Error(err)
		return
	}

	update := database.Db.Where("id = ?", v.UserID).Model(&model.User{}).Updates(map[string]interface{}{"password": hashedPassword})
	if update.Error != nil {
		helper.Logger.Error(update.Error)
		c.JSON(500, gin.H{
			"type":    "email_verification",
			"message": update.Error,
		})
		c.Abort()
		return
	}

}
