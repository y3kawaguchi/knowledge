package main

import (
	"github.com/y3kawaguchi/knowledge/article"
	"github.com/y3kawaguchi/knowledge/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	article := article.New()
	r := gin.Default()
	r.GET("/article", handlers.ArticlesGet(article))
	r.POST("/article", handlers.ArticlePost(article))

	r.Run() // listen and serve on 0.0.0.0:8080

}
