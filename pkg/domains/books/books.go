package books

type Books struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BooksCreate struct {
	Name string `json:"name"`
}

type BooksUpdate struct {
	Name string `json:"name"`
}

type BooksService interface {
	GetBooks(id string) (Books, error)
	CreateBooks(BooksCreate) (Books, error)
	GetAllBooks() ([]Books, error)
	UpdateBooks(id string, input BooksUpdate) (Books, error)
	DeleteBooks(id string) (any, error)
}
