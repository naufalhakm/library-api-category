package params

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BookCategoryRequest struct {
	BookID     uint64 `json:"book_id"`
	CategoryID uint64 `json:"category_id"`
}
