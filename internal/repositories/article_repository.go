package repositories

import (
	"fmt"
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

	fmt.Printf("rows: %#v\n", rows)

	articles := domains.ArticlesNew()
	for rows.Next() {
		item := domains.Article{}
		err := rows.Scan(
			&item.ID,
			&item.AuthorID,
			&item.Title,
			&item.Content,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			// TODO: nilを返すのが適切か考える
			return nil, err
		}
		articles.Add(item)
	}

	fmt.Printf("ArticleRepository.FindAll(): %#v\n", articles)

	return articles, nil
}

// FindByID ...
func (a *ArticleRepository) FindByID(id int64) (*domains.Article, error) {
	db := a.connection.GetDB()

	fmt.Printf("db: %#v\n", db)
	fmt.Printf("id: %#v\n", id)

	query := `SELECT * FROM articles where id = $1`

	item := domains.Article{}
	err := db.QueryRow(query, id).Scan(
		&item.ID,
		&item.AuthorID,
		&item.Title,
		&item.Content,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		// TODO: nilを返すのが適切か考える
		return nil, err
	}

	fmt.Printf("ArticleRepository.FindByID(): %#v\n", item)

	return &item, nil
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

// Update ...
func (a *ArticleRepository) Update(article *domains.Article) (int64, error) {
	now := time.Now()

	_, err := a.connection.GetDB().Exec(
		`UPDATE articles
			SET
				author_id = $1,
				title = $2,
				content = $3,
				updated_at = $4
			WHERE id = $5`,
		article.AuthorID,
		article.Title,
		article.Content,
		now,
		article.ID,
	)
	if err != nil {
		return -1, err
	}

	return 0, err
}
