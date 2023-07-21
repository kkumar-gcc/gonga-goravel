package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	httpswagger "github.com/swaggo/http-swagger"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Web() {
	router := facades.Route()

	router.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})

	// User
	userController := controllers.NewUserController()
	followController := controllers.NewFollowController()
	router.Prefix("users").Group(func(route route.Route) {
		route.Get("/", userController.Index)
		route.Get("/{username}", userController.Show)
		route.Middleware(middleware.Auth()).Put("/{username}", userController.Update)
		route.Middleware(middleware.Auth()).Delete("/{id}", userController.Delete)

		// Follow
		router.Middleware(middleware.Auth()).Post("/follow", followController.Create)
	})

	//Comment
	commentController := controllers.NewCommentController()
	router.Prefix("comments").Group(func(route route.Route) {
		route.Get("/{id}", commentController.Show)
		route.Middleware(middleware.Auth()).Put("/{id}", commentController.Update)
		route.Middleware(middleware.Auth()).Delete("/{id}", commentController.Delete)
	})

	// Post
	postController := controllers.NewPostController()
	router.Prefix("posts").Group(func(route route.Route) {
		route.Get("/", postController.Index)
		route.Get("/{id}", postController.Show)
		route.Middleware(middleware.Auth()).Post("/", postController.Create)
		route.Middleware(middleware.Auth()).Put("/{id}/title", postController.UpdateTitle)
		route.Middleware(middleware.Auth()).Put("/{id}/body", postController.UpdateBody)
		route.Middleware(middleware.Auth()).Put("/{id}/medias", postController.UpdateMedia)
		route.Middleware(middleware.Auth()).Put("/{id}/hashtags", postController.UpdateHashtag)
		route.Middleware(middleware.Auth()).Put("/{id}/settings", postController.UpdatePostSettings)
		route.Middleware(middleware.Auth()).Delete("/{id}", postController.Delete)

		// Comment routes starts with `/posts`
		route.Get("/{id}/comments", commentController.Index)
		route.Middleware(middleware.Auth()).Post("/{id}/comments", commentController.Create)
	})

	// Like
	likeController := controllers.NewLikeController()
	router.Middleware(middleware.Auth()).Post("/likes", likeController.Create)
	router.Middleware(middleware.Auth()).Delete("/likes/{id}", likeController.Delete)

	// Notification
	notificationController := controllers.NewNotificationController()
	router.Middleware(middleware.Auth()).Prefix("notifications").Group(func(route route.Route) {
		route.Get("/", notificationController.Index)
		route.Post("/read_all", notificationController.ReadAll)
		router.Post("/{id}/read", notificationController.Update)
	})

	// Search
	searchController := controllers.NewSearchController()
	router.Get("/search", searchController.Index)

	// Swagger
	swaggerController := controllers.NewSwaggerController()
	router.Get("/swagger", swaggerController.Index)
	router.StaticFile("/swagger.json", "./docs/swagger.json")
	router.Get("/swagger/*any", func(ctx http.Context) {
		handler := httpswagger.Handler(httpswagger.URL("http://localhost:3000/swagger.json"))
		handler(ctx.Response().Writer(), ctx.Request().Origin())
	})

	// Auth Routes
	Auth()
}
