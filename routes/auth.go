package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	auth2 "goravel/app/http/controllers/auth"
	"goravel/app/http/middleware"
)

func Auth() {
	router := facades.Route()

	// Register
	registeredUserController := auth2.NewRegisteredUserController()
	router.Post("/register", registeredUserController.Store)

	// Login
	authenticatedSessionController := auth2.NewAuthenticatedSessionController()
	router.Post("/login", authenticatedSessionController.Store)

	// PasswordReset
	passwordResetLinkController := auth2.NewPasswordResetLinkController()
	router.Post("/forgot-password", passwordResetLinkController.Store)

	// Password
	passwordController := auth2.NewPasswordController()
	router.Post("/reset-password", passwordController.Store)

	// AuthenticatedSession
	router.Middleware(middleware.Auth()).Group(func(route route.Route) {
		// EmailVerificationNotification
		emailVerificationNotificationController := auth2.NewEmailVerificationNotificationController()
		route.Post("/email/verification-notification", emailVerificationNotificationController.Store)

		// VerifyEmail
		verifyEmailController := auth2.NewVerifyEmailController()
		route.Get("/email-verify/{id}/{hash}", verifyEmailController.Index)

		// AuthenticatedSession
		authenticatedSessionController := auth2.NewAuthenticatedSessionController()
		route.Post("/logout", authenticatedSessionController.Destroy)
	})
}
