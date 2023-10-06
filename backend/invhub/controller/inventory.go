package controller

import (
	"invhub/database"
	"invhub/helper"
	"invhub/helper/validation"
	"invhub/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryController struct{}

func (ctrl *InventoryController) Add(c *gin.Context) {
	data := c.MustGet("user").(model.User)

	type rule struct {
		Name     string
		Location string
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

	newInventory := model.Inventory{
		Name:     r.Name,
		Location: &r.Location,
		UserID:   int(data.ID),
	}

	result := database.Db.Create(&newInventory)
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully add a new inventory.",
	})
}
