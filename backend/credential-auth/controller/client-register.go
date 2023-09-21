package controller

import (
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/helper/validation"
	"credential-auth/model"

	"github.com/gin-gonic/gin"
)

type ClientRegisterController struct{}

func (ctrl *ClientRegisterController) Index(c *gin.Context) {
	type rule struct {
		Name   string `json:"name" form:"name" binding:"required"`
		Secret string `json:"secret" form:"secret" binding:"required"`
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

	database.Db.Create(&model.Client{
		Name:   r.Name,
		Secret: r.Secret,
	})

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Successfully register the client",
	})
}
