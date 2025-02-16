package controllers

import (
	"article/config"
	"article/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Posts
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&article)
	c.JSON(http.StatusCreated, article)
}

func GetArticles(c *gin.Context) {
	// Ambil limit & offset dari path parameter (bukan dari query)
	limit, err1 := strconv.Atoi(c.Param("limit"))
	offset, err2 := strconv.Atoi(c.Param("offset"))
	status := c.Query("status") // Ambil status dari query parameter (opsional)

	// Jika limit atau offset tidak valid, beri error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit dan offset harus angka"})
		return
	}

	var articles []models.Posts
	query := config.DB.Limit(limit).Offset(offset)

	// Jika ada filter status, tambahkan ke query
	if status != "" {
		query = query.Where("status = ?", status)
	}

	result := query.Find(&articles)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}



func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article models.Posts

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Posts

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&article)
	c.JSON(http.StatusOK, article)
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Posts{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dihapus"})
}
