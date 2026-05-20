package handlers

import (
	"article-api/app/config"
	"article-api/app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{})
}

func GetArticles(c *gin.Context) {
	var articles []models.Article
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	var limit, offset int
	fmt.Sscanf(limitStr, "%d", &limit)
	fmt.Sscanf(offsetStr, "%d", &offset)

	config.DB.Limit(limit).Offset(offset).Find(&articles)
	c.JSON(http.StatusOK, articles)
}

func GetArticleById(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&article)
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteArticle(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	config.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{})
}
