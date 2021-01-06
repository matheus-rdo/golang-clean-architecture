package presenter

// Book - presenter model (with fewer attributes)
type Book struct {
	ID      int64  `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author"`
}
