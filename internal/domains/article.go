package domains

import "time"

// Article ...
type Article struct {
	ID        int64
	AuthorID  int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Articles ...
type Articles struct {
	Items []Article
}

// ArticlesNew ...
func ArticlesNew() *Articles {
	return &Articles{}
}

// Add ...
func (a *Articles) Add(item Article) {
	a.Items = append(a.Items, item)
}

// GetAll ...
func (a *Articles) GetAll() []Article {
	return a.Items
}

// Change ...
func (a *Article) Change(item Article) *Article {
	return &Article{
		ID:        a.ID,
		AuthorID:  item.AuthorID,
		Title:     item.Title,
		Content:   item.Content,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
