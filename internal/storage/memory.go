package storage

import (
	"fmt"
	"sort"
)

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {
	contact.Id = ms.nextID
	ms.contacts[contact.Id] = contact
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	if len(ms.contacts) == 0 {
		return nil, fmt.Errorf("Contact list is empty")
	}

	contacts := make([]*Contact, 0, len(ms.contacts))
	for _, contact := range ms.contacts {
		contacts = append(contacts, contact)
	}

	sort.Slice(contacts, func(i, j int) bool {
		return contacts[i].Id < contacts[j].Id
	})

	return contacts, nil
}

func (ms *MemoryStore) GetById(id int) (*Contact, error) {
	if _, ok := ms.contacts[id]; !ok {
		return nil, fmt.Errorf("contact with ID %d not found", id)
	}

	contact := *ms.contacts[id]
	return &contact, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	if _, ok := ms.contacts[id]; !ok {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	if newEmail == "" && newName == "" {
		return fmt.Errorf("no update provided for contact with ID %d", id)
	}

	if newName != "" {
		ms.contacts[id].Name = newName
	}
	if newEmail != "" {
		ms.contacts[id].Email = newEmail
	}

	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	delete(ms.contacts, id)
	return nil
}
