package router

import (
	"account-backend/controller"

	"github.com/gin-gonic/gin"
)

var (
	accountController     controller.AccountController     = controller.AccountController{}
	applicationController controller.ApplicationController = controller.ApplicationController{}
)

func Api(r *gin.Engine) {
	r.GET("/profile", accountController.Index)
	r.PUT("/profile", accountController.Update)
	r.PUT("/profile/avatar", accountController.UploadProfilePicture)
	r.GET("/applications", applicationController.Index)
}
