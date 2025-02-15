package routes

import (
	"article/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Perbaiki rute agar tidak konflik
	r.POST("/article", controllers.CreateArticle)
	r.GET("/article/:id", controllers.GetArticleByID)
	r.GET("/article/list/:limit/:offset", controllers.GetArticles)
	r.PUT("/article/:id", controllers.UpdateArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)

	return r
}