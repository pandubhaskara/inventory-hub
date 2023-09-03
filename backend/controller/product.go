package controller

import (
	"invhub/database"
	"invhub/helper"
	"invhub/helper/validation"
	"invhub/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (ctrl *ProductController) Index(c *gin.Context) {
	var p []model.Product

	result := database.Db.Model(model.Product{}).Find(&p)
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
	return
}

func (ctrl *ProductController) Add(c *gin.Context) {
	type rule struct {
		// InventoryID int
		Code     string
		Name     string
		Price    int
		Type     string
		Quantity int
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

	newProduct := model.Product{
		// InventoryID: r.InventoryID,
		Code:     r.Code,
		Name:     r.Name,
		Price:    r.Price,
		Type:     r.Type,
		Quantity: r.Quantity,
	}

	result := database.Db.Create(&newProduct)
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": result.Error,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully add a new product.",
	})
}

func (ctrl *ProductController) Edit(c *gin.Context) {
	id := c.Param("id")

	type rule struct {
		// InventoryID int
		Code     string
		Name     string
		Price    string
		Type     string
		Quantity string
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
	price, err := strconv.Atoi(r.Price)
	quantity, err := strconv.Atoi(r.Quantity)

	newProduct := model.Product{
		// InventoryID: r.InventoryID,
		Code:     r.Code,
		Name:     r.Name,
		Price:    price,
		Type:     r.Type,
		Quantity: quantity,
	}

	result := database.Db.Where("id=?", id).Updates(&newProduct)
	if result.Error != nil {
		helper.Logger.Error(result.Error)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": result.Error,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully edit the product.",
	})
}
