package repositories

import (
	"time"

	"github.com/y3kawaguchi/knowledge/internal/db"
	"github.com/y3kawaguchi/knowledge/internal/domains"
)

// ArticleRepository ...
type ArticleRepository struct {
	connection db.Connection
}

// NewArticleRepository ...
func NewArticleRepository(connection db.Connection) *ArticleRepository {
	return &ArticleRepository{
		connection: connection,
	}
}

// FindAll ...
func (a *ArticleRepository) FindAll() (*domains.Articles, error) {
	db := a.connection.GetDB()

	query := `SELECT * FROM articles`
	rows, err := db.Query(query)
	if err != nil {
		// TODO: nilを返すのが適切か考える
		return nil, err
	}
	defer rows.Close()

	articles := domains.ArticlesNew()
	for rows.Next() {
		item := domains.Article{}
		rows.Scan(
			&item.ID,
			&item.AuthorID,
			&item.Title,
			&item.Content,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		articles.Add(item)
	}

	return articles, nil
}

// Save ...
func (a *ArticleRepository) Save(article *domains.Article) (int64, error) {
	now := time.Now()

	_, err := a.connection.GetDB().Exec(`INSERT INTO articles
		(
			author_id,
			title,
			content,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)`,
		article.AuthorID,
		article.Title,
		article.Content,
		now,
		now,
	)
	if err != nil {
		return -1, err
	}

	return 0, err
}
