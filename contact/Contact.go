package contact

import (
	"errors"
)

type Contact struct {
	Id    int
	Name  string
	Email string
}

func (c *Contact) Add() (*Contact, error) {
	resultCtrl := c.newUser()
	if resultCtrl != "valid" {
		return nil, errors.New("données manquantes pour la création de contact")
	}

	return &Contact{Id: c.Id, Name: c.Name, Email: c.Email}, nil
}

func (c *Contact) newUser() string {
	if c.Name == "" || c.Email == "" {
		return "données manquantes pour l'ajout de contact"
	}
	return "valid"
}

func (c *Contact) Update(name string, mail string) (*Contact, error) {
	if name == "" || mail == "" {
		return nil, errors.New("données manquantes pour la modification de contact")
	}
	c.Name = name
	c.Email = mail
	return c, nil
}
