package controller

import (
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"
	"unicode"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordController struct{}

func (ctrl *ChangePasswordController) Index(c *gin.Context) {
	data := c.MustGet("account").(model.User)

	type rule struct {
		CurrentPassword string `json:"current_password" form:"current_password" binding:"required"`
		NewPassword     string `json:"new_password" form:"new_password" binding:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,min=8"`
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

	// Check current password
	samePassword := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(r.CurrentPassword))
	if samePassword != nil {
		c.JSON(422, gin.H{
			"type":    "change_password",
			"message": "Wrong old password",
		})
		c.Abort()
		return
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range r.NewPassword {
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

	if r.NewPassword != r.ConfirmPassword {
		c.JSON(422, gin.H{
			"type":    "change_password",
			"message": "New password and confirm password unmatched!",
		})
		c.Abort()
		return
	}
	//Hashed new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"type":    "change_password",
			"message": err,
		})
		c.Abort()
		return
	}

	//Update new password
	result := database.Db.Where("id = ?", data.ID).Model(&data).Updates(model.User{
		Password: string(hashedPassword),
	})
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(500, gin.H{
			"type":    "change_password",
			"message": result.Error,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully reset your password",
	})
}
