package controller

import (
	"net/http"
	"strconv"
	"time"

	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"
	"credential-auth/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ForgotPasswordController struct{}

func (ctrl *ForgotPasswordController) Index(c *gin.Context) {
	type rule struct {
		Email    string `json:"email" form:"email" binding:"required"`
		AppBrand string `json:"app_brand"`
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

	var user model.User

	result := database.Db.Where("email = ?", r.Email).First(&user)
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"type":    "forgot_password_email",
			"message": result.Error,
		})
		return
	}

	token := uuid.New()
	expirationTIme, _ := strconv.ParseInt(config.ResetPasswordMailer.ExpirationTime, 10, 64)

	forgotPassword := model.ForgotPassword{
		UserID:    int(user.ID),
		Token:     token.String(),
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add((time.Second * time.Duration(expirationTIme))),
		IsActive:  true,
	}

	forgotPasswordToken := database.Db.Create(&forgotPassword)
	if forgotPasswordToken.Error != nil {
		helper.Logger.Error(forgotPasswordToken.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"type":    "forgot_password_email",
			"message": forgotPasswordToken.Error,
		})
		return
	}

	var email model.Email

	//struct for email

	email = model.Email{
		ReceiverName:    "",
		ReceiverEmail:   user.Email,
		Sender:          config.ResetPasswordMailer.SenderName,
		Subject:         "Reset Password",
		Template:        "forgot-password.html",
		Token:           token.String(),
		ExpirationTime:  config.ResetPasswordMailer.ExpirationTime,
		VerificationURL: config.App.Client + config.ResetPasswordMailer.ResetPasswordUrl,
	}

	//send email module
	service.SendEmail(&email)

	c.JSON(202, gin.H{
		"message": "Instruction to recover your password has been sent to your email.",
	})
}
