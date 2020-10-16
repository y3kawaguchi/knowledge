package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/y3kawaguchi/knowledge/internal/controller/form"
	"github.com/y3kawaguchi/knowledge/internal/domains"

	"github.com/gin-gonic/gin"
)

// Article ...
type Article interface {
	Create(article *domains.Article) (int64, error)
	Get() (*domains.Articles, error)
	Update(article *domains.Article) (int64, error)
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
func (a ArticleAPI) ArticlesGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := a.article.Get()
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

type articlePostRequest struct {
	AuthorID int64  `json:"author_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// ArticlePost ...
func (a ArticleAPI) ArticlePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form.Article
		if err := c.ShouldBindJSON(&f); err != nil {
			c.Error(err).SetMeta(http.StatusUnprocessableEntity)
			return
		}
		article := f.BuildDomain()
		_, err := a.article.Create(article)
		if err != nil {
			fmt.Printf("error: %#v\n", err)
			c.Error(err).SetMeta(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}

// ArticlePut ...
func (a ArticleAPI) ArticlePut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form.Article
		if err := c.ShouldBindJSON(&f); err != nil {
			c.Error(err).SetMeta(http.StatusUnprocessableEntity)
			return
		}

		id, err := strconv.Atoi(c.Param("articles_id"))
		if err != nil {
			c.Error(err).SetMeta(http.StatusNotFound)
			return
		}

		article := f.BuildDomain()
		article.ID = int64(id)

		if _, err := a.article.Update(article); err != nil {
			fmt.Printf("error: %#v\n", err)
			c.Error(err).SetMeta(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}
