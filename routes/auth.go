package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	Auth2 "goravel/app/http/controllers/Auth"
	"goravel/app/http/middleware"
)

func Auth() {
	router := facades.Route()

	// Register
	registeredUserController := Auth2.NewRegisteredUserController()
	router.Post("/register", registeredUserController.Store)

	// Login
	authenticatedSessionController := Auth2.NewAuthenticatedSessionController()
	router.Post("/login", authenticatedSessionController.Store)

	// PasswordReset
	passwordResetLinkController := Auth2.NewPasswordResetLinkController()
	router.Post("/forgot-password", passwordResetLinkController.Store)

	// Password
	passwordController := Auth2.NewPasswordController()
	router.Post("/reset-password", passwordController.Store)

	// AuthenticatedSession
	router.Middleware(middleware.Auth()).Group(func(route route.Route) {
		// EmailVerificationNotification
		emailVerificationNotificationController := Auth2.NewEmailVerificationNotificationController()
		route.Post("/email/verification-notification", emailVerificationNotificationController.Store)

		// VerifyEmail
		verifyEmailController := Auth2.NewVerifyEmailController()
		route.Get("/email-verify/{id}/{hash}", verifyEmailController.Index)

		// AuthenticatedSession
		authenticatedSessionController := Auth2.NewAuthenticatedSessionController()
		route.Post("/logout", authenticatedSessionController.Destroy)
	})
}
