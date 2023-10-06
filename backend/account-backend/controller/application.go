package controller

import (
	"account-backend/database"
	"account-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationController struct{}

func (ctlr *ApplicationController) Index(c *gin.Context) {
	data := c.MustGet("account").(model.Account)

	// Get the applications
	var account model.Account
	database.Db.Preload("Applications").Where("id = ? ", data.ID).First(&account)

	// Return a success message to the client
	c.JSON(http.StatusOK, gin.H{
		"data": account.Applications,
	})
}
