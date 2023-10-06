package controller

import (
	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"
	"credential-auth/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterController struct{}

func (ctrl *RegisterController) ResendVerification(c *gin.Context) {
	//CHANGES USING ACCESS TOKEN FROM LOGIN
	type rule struct {
		Email string `json:"email" form:"email" binding:"required"`
	}

	var r rule

	err := c.ShouldBind(&r)
	if err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	//create data to database
	database.Db.Transaction(func(tx *gorm.DB) error {
		verificationToken := uuid.New()

		var u model.User
		// save new user data
		result := tx.Where("email = ?", r.Email).First(&u)
		if result.Error != nil {
			helper.Logger.Error(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": result.Error,
			})
			c.Abort()
			return result.Error
		}

		tx.Model(&model.VerificationToken{}).Where("user_id = ?", u.ID).Update("is_active", "false")
		// save verification token
		verificationTokenDuration, _ := strconv.ParseInt(config.RegistrationMailer.ExpirationTime, 10, 64)

		newVerificationToken := model.VerificationToken{
			UserID:    int(u.ID),
			Token:     verificationToken.String(),
			IssuedAt:  time.Now(),
			ExpiredAt: time.Now().Add((time.Second * time.Duration(verificationTokenDuration))),
			IsActive:  true,
		}

		saveVerificationToken := tx.Create(&newVerificationToken)
		if saveVerificationToken.Error != nil {
			helper.Logger.Error(saveVerificationToken.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"type":    "registration_account",
				"message": saveVerificationToken.Error,
			})
			c.Abort()
			return saveVerificationToken.Error
		}

		// struct for email
		email := model.Email{
			ReceiverName:    "",
			ReceiverEmail:   u.Email,
			Sender:          config.RegistrationMailer.SenderName,
			Subject:         "Account Verification",
			Template:        "verification-email.html",
			Token:           verificationToken.String(),
			ExpirationTime:  config.RegistrationMailer.ExpirationTime,
			VerificationURL: config.App.Client + config.RegistrationMailer.VerificationUrl,
		}

		// send email module
		go service.SendEmail(&email)

		c.JSON(200, gin.H{
			"message": "Successfully resend email.",
		})
		return nil
	})
}

func (ctrl *RegisterController) NewUser(c *gin.Context) {
	type rule struct {
		Name     string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var r rule

	err := c.ShouldBind(&r)
	if err != nil {
		helper.Logger.Error(err)
		code, e := validation.HttpRequestValidationErrorJson(err)
		c.JSON(code, e)
		c.Abort()
		return
	}

	// Store new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.Logger.Error(err)
		return
	}

	result := database.Db.Create(&model.User{
		Email:    r.Email,
		Username: r.Username,
		Password: string(hashedPassword),
	})
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully register.",
	})
	return
}
