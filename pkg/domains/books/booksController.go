package books

import (
	"github.com/go-fuego/fuego"
)

type BooksResources struct {
	// TODO add resources
	BooksService BooksService
}

func (rs BooksResources) Routes(s *fuego.Server) {
	booksGroup := fuego.Group(s, "/books")

	fuego.Get(booksGroup, "/", rs.getAllBooks)
	fuego.Post(booksGroup, "/", rs.postBooks)

	fuego.Get(booksGroup, "/{id}", rs.getBooks)
	fuego.Put(booksGroup, "/{id}", rs.putBooks)
	fuego.Delete(booksGroup, "/{id}", rs.deleteBooks)
}

func (rs BooksResources) getAllBooks(c fuego.ContextNoBody) ([]Books, error) {
	return rs.BooksService.GetAllBooks()
}

func (rs BooksResources) postBooks(c fuego.ContextWithBody[BooksCreate]) (Books, error) {
	body, err := c.Body()
	if err != nil {
		return Books{}, err
	}

	return rs.BooksService.CreateBooks(body)
}

func (rs BooksResources) getBooks(c fuego.ContextNoBody) (Books, error) {
	id := c.PathParam("id")

	return rs.BooksService.GetBooks(id)
}

func (rs BooksResources) putBooks(c fuego.ContextWithBody[BooksUpdate]) (Books, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return Books{}, err
	}

	return rs.BooksService.UpdateBooks(id, body)
}

func (rs BooksResources) deleteBooks(c fuego.ContextNoBody) (any, error) {
	return rs.BooksService.DeleteBooks(c.PathParam("id"))
}
