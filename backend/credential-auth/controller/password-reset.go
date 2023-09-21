package controller

import (
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"
	"errors"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PasswordResetController struct{}

func (ctrl *PasswordResetController) Index(c *gin.Context) {
	token := c.Param("token")

	var f model.ForgotPassword

	result := database.Db.Where("token=? AND is_active=? AND expired_at>= NOW()", token, true).First(&f)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"type":    "password_reset",
				"message": "Invalid token.",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(500, gin.H{
				"type":    "password_reset",
				"message": result.Error,
			})
			c.Abort()
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Valid token.",
	})
}

func (ctrl *PasswordResetController) Store(c *gin.Context) {
	type rule struct {
		Token                string `json:"token" form:"token" binding:"required"`
		Password             string `json:"password" form:"token" password:"required"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"required"`
	}

	var r rule
	var f model.ForgotPassword

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	result := database.Db.Where("token=? AND is_active=? AND expired_at>= NOW()", r.Token, true).First(&f)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"type":    "password_reset",
				"message": "Invalid token.",
			})
			c.Abort()
			return

		} else {
			helper.Logger.Error(result.Error)
			c.JSON(401, gin.H{
				"type":    "password_reset",
				"message": result.Error,
			})
			c.AbortWithStatus(500)
			return
		}
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range r.Password {
		if unicode.IsUpper(v) {
			hasUpper = true
		} else if unicode.IsLower(v) {
			hasLower = true
		} else if unicode.IsNumber(v) {
			hasNumber = true
		} else if unicode.IsGraphic(v) {
			hasSymbol = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSymbol {
		c.JSON(422, gin.H{
			"type":    "change_password",
			"message": "Password must be at least 8 characters long and consist of 1 number, 1 uppercase letter, 1 lowercase letter and 1 special character.",
		})
		c.Abort()
		return
	}

	if r.Password != r.PasswordConfirmation {
		c.JSON(422, gin.H{
			"type":    "change_password",
			"message": "New password and confirm password unmatched!",
		})
		c.Abort()
		return
	}

	// Hashed new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.Logger.Error(err)
		c.JSON(500, gin.H{
			"type":    "change_password",
			"message": err,
		})
		c.Abort()
		return
	}

	//Update new password
	update := database.Db.Where("id = ?", f.UserID).Model(&model.User{}).Update("password", string(hashedPassword))
	if update.Error != nil {
		helper.Logger.Error(update.Error)
		c.JSON(500, gin.H{
			"type":    "change_password",
			"message": update.Error,
		})
		c.Abort()
		return
	}

	t := time.Now()
	f.IsActive = false
	f.UsedAt = &t
	deactiveToken := database.Db.Save(&f)

	if deactiveToken.Error != nil {
		if errors.Is(deactiveToken.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"type":    "password_reset",
				"message": "Invalid token.",
			})
			c.Abort()
			return

		} else {
			helper.Logger.Error(deactiveToken.Error)
			c.JSON(500, gin.H{
				"type":    "change_password",
				"message": deactiveToken.Error,
			})
			c.Abort()
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Password has been successfuly changed.",
	})
}
