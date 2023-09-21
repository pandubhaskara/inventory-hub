package router

import (
	"credential-auth/controller"
	"credential-auth/middleware"

	"github.com/gin-gonic/gin"
)

var (
	loginController             controller.LoginController             = controller.LoginController{}
	introspectController        controller.IntrospectController        = controller.IntrospectController{}
	changeEmailController       controller.ChangeEmailController       = controller.ChangeEmailController{}
	forgotPasswordController    controller.ForgotPasswordController    = controller.ForgotPasswordController{}
	changePasswordController    controller.ChangePasswordController    = controller.ChangePasswordController{}
	passwordResetController     controller.PasswordResetController     = controller.PasswordResetController{}
	registerController          controller.RegisterController          = controller.RegisterController{}
	emailVerificationController controller.EmailVerificationController = controller.EmailVerificationController{}
)

func Api(r *gin.Engine) {
	internal := r.Group("_i")
	{
		internal.POST("/introspect", introspectController.Index)
	}

	r.POST("/token", loginController.Index)
	r.GET("/password-reset/:token", passwordResetController.Index)
	r.GET("/verify-account/:token", emailVerificationController.Index)
	r.POST("/forgot-password", forgotPasswordController.Index)
	r.POST("/password-reset", passwordResetController.Store)
	r.POST("/resend-verification", registerController.ResendVerification)
	r.POST("/verify-account", emailVerificationController.Store)
	r.POST("/set-password", emailVerificationController.SetPassword)

	user := r.Group("user", middleware.Auth)
	{
		user.POST("/change-password", changePasswordController.Index)
		user.POST("/change-email", changeEmailController.Index)
	}
}
