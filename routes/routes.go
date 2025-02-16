package routes

import (
	"article/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	articleRoutes := r.Group("/article")
	{
		articleRoutes.POST("/", controllers.CreateArticle)
		articleRoutes.GET("/list/:limit/:offset", controllers.GetArticles)
		articleRoutes.GET("/:id", controllers.GetArticleByID)
		articleRoutes.PUT("/:id", controllers.UpdateArticle)
		articleRoutes.DELETE("/:id", controllers.DeleteArticle)
	}
}
