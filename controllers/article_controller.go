package controllers

import (
	"article/config"
	"article/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Membuat artikel baru
func CreateArticle(c *gin.Context) {
	var article models.Posts
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input
	if err := validate.Struct(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}

// Menampilkan semua artikel dengan pagination
func GetArticles(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))

	var articles []models.Posts
	config.DB.Limit(limit).Offset(offset).Find(&articles)

	c.JSON(http.StatusOK, articles)
}

// Menampilkan satu artikel berdasarkan ID
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article models.Posts
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

// Mengupdate artikel berdasarkan ID
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Posts

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully"})
}

// Menghapus artikel berdasarkan ID
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Posts{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}