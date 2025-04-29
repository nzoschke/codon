package contacts

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-fuego/fuego"
)

type Service struct {
	db map[string]Contact
	mu sync.RWMutex
}

var _ Contacts = &Service{}

func New() Contacts {
	return &Service{
		db: make(map[string]Contact),
	}
}

func (s *Service) Get(id string) (Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	contact, exists := s.db[id]
	if !exists {
		return Contact{}, fuego.NotFoundError{Title: "Contact not found with id " + id}
	}

	return contact, nil
}

func (s *Service) Create(input ContactCreateIn) (Contact, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	contact := Contact{
		ID:   id,
		Name: input.Name,
	}

	s.db[id] = contact
	return contact, nil
}

func (s *Service) List() ([]Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	contacts := make([]Contact, 0, len(s.db))
	for _, books := range s.db {
		contacts = append(contacts, books)
	}

	return contacts, nil
}

func (s *Service) Update(id string, input ContactUpdateIn) (Contact, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact, exists := s.db[id]
	if !exists {
		return Contact{}, fuego.NotFoundError{Title: "Contact not found with id " + id}
	}

	if input.Name != "" {
		contact.Name = input.Name
	}

	s.db[id] = contact
	return contact, nil
}

func (s *Service) Delete(id string) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.db[id]
	if !exists {
		return nil, fuego.NotFoundError{Title: "Contact not found with id " + id}
	}

	delete(s.db, id)
	return "deleted", nil
}
