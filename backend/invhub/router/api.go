package router

import (
	"invhub/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	UserController    controller.UserController    = controller.UserController{}
	ProductController controller.ProductController = controller.ProductController{}
)

func Api(r *gin.Engine) {
	api := r.Group("api")
	api.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "public router is working...",
		})
	})
	auth := api.Group("auth")
	{
		auth.POST("/register", UserController.Register)
		auth.POST("/login", UserController.Login)
	}
	{
		v1 := api.Group("v1")
		{
			v1.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "authenticated router is working...",
				})
			})

			product := v1.Group("products")
			{
				product.GET("", ProductController.Index)
				product.POST("", ProductController.Add)
				product.PUT("/edit/:id", ProductController.Edit)
			}
		}
	}

}
