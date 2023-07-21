package routes

import (
	"github.com/goravel/framework/facades"
	Auth2 "goravel/app/http/controllers/Auth"
	"goravel/app/http/middleware"
)

func Auth() {
	router := facades.Route()

	// Login\Logout
	loginController := Auth2.NewLoginController()
	router.Post("/login", loginController.Create)
	router.Middleware(middleware.Auth()).Post("/logout", loginController.Delete)

	// Register
	registerController := Auth2.NewRegisterController()
	router.Post("/register", registerController.Create)

	// PasswordReset
	passwordResetLinkController := Auth2.NewPasswordResetLinkController()
	router.Post("/forgot-password", passwordResetLinkController.Create)

	// NewPassword
	newPasswordController := Auth2.NewNewPasswordController()
	router.Post("/reset-password", newPasswordController.Create)
}
