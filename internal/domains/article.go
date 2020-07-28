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
func (r *Articles) Add(item Article) {
	r.Items = append(r.Items, item)
}

// GetAll ...
func (r *Articles) GetAll() []Article {
	return r.Items
}
