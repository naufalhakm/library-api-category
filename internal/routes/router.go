package routes

import (
	"fmt"
	"library-api-category/internal/factory"
	"library-api-category/internal/grpc/client"
	"library-api-category/internal/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(provider *factory.Provider, authClient *client.AuthClient) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), CORS())

	router.GET("/", func(ctx *gin.Context) {
		currentYear := time.Now().Year()
		message := fmt.Sprintf("Library API Category %d", currentYear)

		ctx.JSON(http.StatusOK, message)
	})

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			auth := v1.Use(middleware.CheckAuth(authClient))
			auth.GET("/categories", provider.CategoryProvider.GetAllCategories)
			auth.GET("/categories/:id", provider.CategoryProvider.GetDetailCategory)
			auth.GET("/categories/books/:id", provider.CategoryProvider.ListCategoryOfBook)

			admin := v1.Use(middleware.CheckAuthIsAdminOrAuthor(authClient))
			admin.POST("/categories", provider.CategoryProvider.CreateCategory)
			admin.PUT("/categories/:id", provider.CategoryProvider.UpdateCategory)
			admin.DELETE("/categories/:id", provider.CategoryProvider.DeleteCategory)
			admin.POST("/categories/books", provider.CategoryProvider.AddBookCategory)
		}
	}

	return router
}

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, accept, access-control-allow-origin, access-control-allow-headers")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
