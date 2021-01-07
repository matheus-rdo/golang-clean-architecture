package presenter

// Book - presenter model (with fewer attributes)
type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author"`
}
