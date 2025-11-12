package storage

import "fmt"

type Contact struct {
	Id    int
	Name  string
	Email string
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetById(id int) (*Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

var ErrContactNotFound = func(id int) error { return fmt.Errorf("Contact with ID %d not found", id) }
