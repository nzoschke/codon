package contacts

import (
	"github.com/go-fuego/fuego"
)

type Resources struct {
	Contacts Contacts
}

func (rs Resources) Routes(s *fuego.Server) {
	g := fuego.Group(s, "/contacts")

	fuego.Get(g, "/", rs.list)
	fuego.Post(g, "/", rs.create)
	fuego.Get(g, "/{id}", rs.get)
	fuego.Put(g, "/{id}", rs.update)
	fuego.Delete(g, "/{id}", rs.delete)
}

func (rs Resources) create(c fuego.ContextWithBody[ContactCreateIn]) (Contact, error) {
	body, err := c.Body()
	if err != nil {
		return Contact{}, err
	}

	return rs.Contacts.Create(body)
}

func (rs Resources) delete(c fuego.ContextNoBody) (any, error) {
	return rs.Contacts.Delete(c.PathParam("id"))
}

func (rs Resources) get(c fuego.ContextNoBody) (Contact, error) {
	return rs.Contacts.Get(c.PathParam("id"))
}

func (rs Resources) list(c fuego.ContextNoBody) ([]Contact, error) {
	return rs.Contacts.List()
}

func (rs Resources) update(c fuego.ContextWithBody[ContactUpdateIn]) (Contact, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return Contact{}, err
	}

	return rs.Contacts.Update(id, body)
}
