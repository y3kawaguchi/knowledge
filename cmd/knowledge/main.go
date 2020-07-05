package main

import (
	"github.com/y3kawaguchi/knowledge/internal/db"
	"github.com/y3kawaguchi/knowledge/internal/handlers"
	"github.com/y3kawaguchi/knowledge/internal/repositories"
	"github.com/y3kawaguchi/knowledge/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: plan to remove
	//articles := domains.ArticlesNew()

	// PostgreSQL setup
	dbConfig, err := db.GetPostgreSQLConfigFromEnv()
	if err != nil {
		// TODO: plan to implement logging
	}
	dbConnection, err := db.ConnectPostgreSQL(dbConfig)
	if err != nil {
		// TODO: plan to implement logging
	}

	// inject dbConnection to repository
	articleRepository := repositories.NewArticleRepository(dbConnection)

	// inject articleRepository to usecase
	articleUsecase := usecases.NewArticleUsecase(articleRepository)

	articleAPI := handlers.NewArticleAPI(articleUsecase)

	r := gin.Default()
	// r.GET("/article", handlers.ArticlesGet(articles))
	r.GET("/article", articleAPI.ArticlesGet())
	//	r.POST("/article", handlers.ArticlePost(articles))

	// listen and serve on 0.0.0.0:8080
	r.Run()
}

// func main() {
// 	// TODO: plan to remove
// 	article := domains.ArticlesNew()

// 	container := dig.New()

// 	// PostgreSQL setup
// 	dbConfig, err := database.GetPostgreSQLConfigFromEnv()
// 	if err != nil {
// 		// TODO: plan to implement logging
// 	}
// 	dbConnection, err := database.ConnectPostgreSQL(dbConfig)
// 	if err != nil {
// 		// TODO: plan to implement logging
// 	}
// 	// container.Invoke()した時に実行して欲しいfunc(dbConnection.GetDB)を登録
// 	err = container.Provide(dbConnection.GetDB)
// 	if err != nil {
// 		// TODO: plan to implement logging
// 	}

// 	r := gin.Default()
// 	r.GET("/article", handlers.ArticlesGet(article))
// 	r.POST("/article", handlers.ArticlePost(article))

// 	// listen and serve on 0.0.0.0:8080
// 	r.Run()
// }
