package main

import (
	"net/http"

	"github.com/rfinochi/golang-workshop-todo/docs"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) initRouter() {
	app.router = gin.Default()
}

func (app *application) addApiRoutes() {
	if app.router != nil {
		app.router.Use(static.Serve("/", static.LocalFile("./ui/html", true)))

		api := app.router.Group("/api")
		api.GET("/", app.getItemsEndpoint)
		api.GET("/:id", app.getItemEndpoint)
		api.POST("/", app.postItemEndpoint)
		api.PUT("/", app.putItemEndpoint)
		api.PATCH("/:id", app.updateItemEndpoint)
		api.DELETE("/:id", app.deleteItemEndpoint)
	}
}

func (app *application) addSwaggerRoutes() {
	if app.router != nil {
		docs.SwaggerInfo.Schemes = []string{"https", "http"}
		app.router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		app.router.GET("/api-docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "./api-docs/index.html")
		})
	}
}
