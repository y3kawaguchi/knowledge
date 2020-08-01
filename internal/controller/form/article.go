package form

import (
	"github.com/y3kawaguchi/knowledge/internal/domains"
)

// Article ...
type Article struct {
	AuthorID int64  `json:"author_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// BuildDomain ...
func (f *Article) BuildDomain() *domains.Article {
	return &domains.Article{
		AuthorID: f.AuthorID,
		Title:    f.Title,
		Content:  f.Content,
	}
}
