package controller

import (
	"errors"
	"invhub/database"
	"invhub/helper"
	"invhub/helper/validation"
	"invhub/model"
	"invhub/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct{}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (ctrl *UserController) Register(c *gin.Context) {
	type rule struct {
		Username string `json:"username" binding:"required"`
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

	hashedPassword, err := HashPassword(r.Password)
	if err != nil {
		helper.Logger.Error(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}

	newUser := model.User{
		Username: r.Username,
		Password: hashedPassword,
	}

	result := database.Db.Create(&newUser)
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": result.Error,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully register.",
	})
}

func (ctrl *UserController) Login(c *gin.Context) {
	type rule struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
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

	var u model.User

	result := database.Db.Where("username = ?", r.Username).First(&u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Wrong username/password",
			})
			c.Abort()
			return
		} else {
			helper.Logger.Error(result.Error)
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": result.Error,
			})
			c.AbortWithStatus(500)
			return
		}
	}

	//comparing the password with the hash
	match := CheckPasswordHash(r.Password, u.Password)
	if !match {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "Wrong email/password",
		})
		c.Abort()
		return
	}

	//create token
	token, err := service.GenerateToken(u)
	if err != nil {
		helper.Logger.Error(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, token)
}
