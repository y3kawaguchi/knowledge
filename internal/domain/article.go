package domain

// Article ...
type Article struct {
	AuthorID int64
	Title    string
	Content  string
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
