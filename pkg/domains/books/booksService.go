package books

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-fuego/fuego"
)

type BooksServiceImpl struct {
	booksRepository map[string]Books
	mu                  sync.RWMutex
}

var _ BooksService = &BooksServiceImpl{}

func NewBooksService() BooksService {
	return &BooksServiceImpl{
		booksRepository: make(map[string]Books),
	}
}

func (bs *BooksServiceImpl) GetBooks(id string) (Books, error) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	books, exists := bs.booksRepository[id]
	if !exists {
		return Books{}, fuego.NotFoundError{Title: "Books not found with id " + id}
	}

	return books, nil
}

func (bs *BooksServiceImpl) CreateBooks(input BooksCreate) (Books, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	books := Books{
		ID:   id,
		Name: input.Name,
	}

	bs.booksRepository[id] = books
	return books, nil
}

func (bs *BooksServiceImpl) GetAllBooks() ([]Books, error) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	allBooks := make([]Books, 0, len(bs.booksRepository))
	for _, books := range bs.booksRepository {
		allBooks = append(allBooks, books)
	}

	return allBooks, nil
}

func (bs *BooksServiceImpl) UpdateBooks(id string, input BooksUpdate) (Books, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	books, exists := bs.booksRepository[id]
	if !exists {
		return Books{}, fuego.NotFoundError{Title: "Books not found with id " + id}
	}

	if input.Name != "" {
		books.Name = input.Name
	}

	bs.booksRepository[id] = books
	return books, nil
}

func (bs *BooksServiceImpl) DeleteBooks(id string) (any, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	_, exists := bs.booksRepository[id]
	if !exists {
		return nil, fuego.NotFoundError{Title: "Books not found with id " + id}
	}

	delete(bs.booksRepository, id)
	return "deleted", nil
}
