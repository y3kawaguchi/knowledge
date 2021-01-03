package main

import (
	"time"

	"github.com/y3kawaguchi/knowledge/internal/controller/handlers"
	"github.com/y3kawaguchi/knowledge/internal/db"
	"github.com/y3kawaguchi/knowledge/internal/repositories"
	"github.com/y3kawaguchi/knowledge/internal/usecases"

	"github.com/gin-gonic/gin"
)

const (
	location = "Asia/Tokyo"
	offset   = 9 * 60 * 60
)

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, offset)
	}
	time.Local = loc
}

func main() {
	// PostgreSQL setup
	dbConfig, err := db.GetPostgreSQLConfigFromEnv()
	if err != nil {
		// TODO: plan to implement logging
	}
	dbConnection, err := db.ConnectPostgreSQL(dbConfig)
	if err != nil {
		// TODO: plan to implement logging
	}
	defer dbConnection.Close()

	// inject dbConnection to repository
	articleRepository := repositories.NewArticleRepository(dbConnection)

	// inject articleRepository to usecase
	articleUsecase := usecases.NewArticleUsecase(articleRepository)

	// inject articleUsecase to handler
	articleAPI := handlers.NewArticleAPI(articleUsecase)

	r := gin.Default()
	r.GET("/articles", articleAPI.ArticlesGet())
	r.POST("/articles", articleAPI.ArticlePost())
	r.PUT("/articles/:articles_id", articleAPI.ArticlePut())

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
