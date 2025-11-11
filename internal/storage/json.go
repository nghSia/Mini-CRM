package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/nghSia/Mini-CRM/internal/config"
)

type JSONStore struct {
	contacts map[int]*Contact
	nextID   int
	path     string
}

func GetContactsFilePath() string {
	workingDir, err := os.Getwd()
	if err != nil {
		return "contacts.json"
	}
	return filepath.Join(workingDir, "contacts.json")
}

func NewJsonStore() *JSONStore {
	js := &JSONStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
		path:     GetContactsFilePath(),
	}
	_ = js.loadFromFile()
	return js
}

func (js *JSONStore) convertToInputTargets() []config.InputTarget {
	var targets []config.InputTarget
	for _, c := range js.contacts {
		target := config.InputTarget{
			Id:    c.Id,
			Name:  c.Name,
			Email: c.Email,
		}
		targets = append(targets, target)
	}
	return targets
}

func (js *JSONStore) loadFromFile() error {
	targets, err := config.LoadTargetsFromFile(js.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to load contacts from file: %w", err)
	}

	js.contacts = make(map[int]*Contact)
	js.nextID = 1

	for _, t := range targets {
		contact := &Contact{
			Id:    t.Id,
			Name:  t.Name,
			Email: t.Email,
		}
		js.contacts[contact.Id] = contact

		if contact.Id >= js.nextID {
			js.nextID = contact.Id + 1
		}
	}

	return nil
}

func (js *JSONStore) Add(contact *Contact) error {
	contact.Id = js.nextID
	js.contacts[contact.Id] = contact
	js.nextID++

	targets := js.convertToInputTargets()

	err := config.SaveTargetsToFile(targets, js.path)
	if err != nil {
		return fmt.Errorf("failed to save contacts to file: %w", err)
	}

	return nil
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	if err := js.loadFromFile(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	if len(js.contacts) == 0 {
		return nil, fmt.Errorf("Contact list is empty")
	}

	contacts := make([]*Contact, 0, len(js.contacts))
	for _, contact := range js.contacts {
		contacts = append(contacts, contact)
	}

	sort.Slice(contacts, func(i, j int) bool {
		return contacts[i].Id < contacts[j].Id
	})

	return contacts, nil
}

func (js *JSONStore) GetById(id int) (*Contact, error) {
	if err := js.loadFromFile(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	if _, ok := js.contacts[id]; !ok {
		return nil, fmt.Errorf("contact with ID %d not found", id)
	}

	contact := *js.contacts[id]
	return &contact, nil
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	if _, ok := js.contacts[id]; !ok {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	if newEmail == "" && newName == "" {
		return fmt.Errorf("no update provided for contact with ID %d", id)
	}

	if newName != "" {
		js.contacts[id].Name = newName
	}
	if newEmail != "" {
		js.contacts[id].Email = newEmail
	}

	targets := js.convertToInputTargets()

	err := config.SaveTargetsToFile(targets, js.path)
	if err != nil {
		return fmt.Errorf("failed to save contacts to file: %w", err)
	}

	return nil
}

func (js *JSONStore) Delete(id int) error {
	if _, ok := js.contacts[id]; !ok {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	delete(js.contacts, id)

	targets := js.convertToInputTargets()

	err := config.SaveTargetsToFile(targets, js.path)
	if err != nil {
		return fmt.Errorf("failed to save contacts to file: %w", err)
	}

	return nil
}
