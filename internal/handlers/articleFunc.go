package handlers

import (
	"net/http"

	"github.com/y3kawaguchi/knowledge/article"

	"github.com/gin-gonic/gin"
)

// ArticlesGet ...
func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := articles.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

type articlePostRequest struct {
	AuthorID int64  `json:"author_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// ArticlePost ...
func ArticlePost(post *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := articlePostRequest{}
		c.Bind(&requestBody)

		item := article.Item{
			AuthorID: requestBody.AuthorID,
			Title:    requestBody.Title,
			Content:  requestBody.Content,
		}
		post.Add(item)

		c.Status(http.StatusNoContent)
	}
}
