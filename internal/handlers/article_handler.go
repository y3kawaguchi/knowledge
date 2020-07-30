package handlers

import (
	"fmt"
	"net/http"

	"github.com/y3kawaguchi/knowledge/internal/domains"

	"github.com/gin-gonic/gin"
)

// Article ...
type Article interface {
	Create(article *domains.Article) (int64, error)
	// Get() (*domains.Articles, error)
}

// ArticleAPI ...
type ArticleAPI struct {
	article Article
}

// NewArticleAPI ...
func NewArticleAPI(article Article) *ArticleAPI {
	return &ArticleAPI{
		article: article,
	}
}

// ArticlesGet ...
// func (a ArticleAPI) ArticlesGet() gin.HandlerFunc {
// 	// func ArticlesGet(articles *domains.Articles) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// result := articles.GetAll()

// 		// TODO: 変数化する
// 		result, err := a.article.Get()
// 		if err != nil {
// 			c.Error(err).SetMeta(http.StatusInternalServerError)
// 			return
// 		}
// 		c.JSON(http.StatusOK, result)
// 	}
// }

type articlePostRequest struct {
	AuthorID int64  `json:"author_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// ArticlePost ...
func (a ArticleAPI) ArticlePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := articlePostRequest{}
		c.Bind(&requestBody)

		item := &domains.Article{
			AuthorID: requestBody.AuthorID,
			Title:    requestBody.Title,
			Content:  requestBody.Content,
		}

		_, err := a.article.Create(item)

		if err != nil {
			fmt.Printf("error: %#v\n", err)
			c.Error(err).SetMeta(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}
