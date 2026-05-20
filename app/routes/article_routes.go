package routes

import (
	"article-api/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	articleRoutes := r.Group("/article")
	{
		articleRoutes.POST("/", handlers.CreateArticle)
		articleRoutes.GET("/", handlers.GetArticles)
		articleRoutes.GET("/:id", handlers.GetArticleById)
		articleRoutes.PUT("/:id", handlers.UpdateArticle)
		articleRoutes.DELETE("/:id", handlers.DeleteArticle)
	}
}
