package handlers

import (
	"article-api/app/config"
	"article-api/app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ArticleInput defines the structure for incoming request data and its validation rules
type ArticleInput struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status" binding:"required,oneof=publish draft thrash"`
}

func CreateArticle(c *gin.Context) {
	var input ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   input.Status,
	}

	config.DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{})
}

func GetArticles(c *gin.Context) {
	var articles []models.Article
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	status := c.Query("status")

	var limit, offset int
	fmt.Sscanf(limitStr, "%d", &limit)
	fmt.Sscanf(offsetStr, "%d", &offset)

	query := config.DB

	if status == "thrash" {
		query = query.Unscoped().Where("deleted_at IS NOT NULL")
	} else if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Limit(limit).Offset(offset).Find(&articles)
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

	var input ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&article).Updates(models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   input.Status,
	})

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
