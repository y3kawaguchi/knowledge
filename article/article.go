package article

// Item ...
type Item struct {
	AuthorID int64
	Title    string
	Content  string
}

// Articles ...
type Articles struct {
	Items []Item
}

// New ...
func New() *Articles {
	return &Articles{}
}

// Add ...
func (r *Articles) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll ...
func (r *Articles) GetAll() []Item {
	return r.Items
}
