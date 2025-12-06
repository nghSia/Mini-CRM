package storage

import (
	"fmt"
)

type Contact struct {
	Id    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(100);unique;not null"`
	Email string `gorm:"type:varchar(100);unique;not null"`
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetById(id int) (*Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

var ErrContactNotFound = func(id int) error { return fmt.Errorf("Contact with ID %d not found", id) }
