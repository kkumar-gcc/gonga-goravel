package middleware

import (
	"errors"
	"strings"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

// Auth is a middleware function that checks if a user is authenticated.
// If the user is not authenticated, it returns an error response with status code 401.
// If the user is authenticated, it calls the next middleware/handler in the chain.
func Auth() http.Middleware {
	return func(ctx http.Context) {
		token, err := extractToken(ctx)
		if err != nil {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if _, err := facades.Auth().Parse(ctx, token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = facades.Auth().Refresh(ctx)
				if err != nil {
					// Refresh time exceeded
					ctx.Request().AbortWithStatus(http.StatusUnauthorized)
					return
				}

				token = "Bearer " + token
			} else {
				ctx.Request().AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		// You can get User in DB and set it to ctx
		var user models.User
		if err := facades.Auth().User(ctx, &user); err != nil {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.WithValue("user", user)
		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}
}

// extractToken retrieves the token from the Authorization header.
func extractToken(ctx http.Context) (string, error) {
	authorizationHeader := ctx.Request().Header("Authorization", "")
	if authorizationHeader == "" {
		return "", errors.New("missing authorization header")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid or unsupported token type")
	}

	return headerParts[1], nil
}
