package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore() *GORMStore {
	var err error

	// Create database directory if it doesn't exist
	dbDir := "database"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create database directory: %v", err)
	}

	dbPath := filepath.Join(dbDir, "contacts.db")
	log.Printf("🔄 Trying to connect to the database %s", dbPath)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: '%s' : %v", dbPath, err)
	}

	err = db.AutoMigrate(&Contact{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	log.Printf("✅ Successfully connected to the database %s", dbPath)
	return &GORMStore{db: db}
}

func (gs *GORMStore) Add(contact *Contact) error {
	result := gs.db.Create(contact)
	if result.Error != nil {
		return fmt.Errorf("failed to add contact: %w", result.Error)
	}
	return nil
}

func (gs *GORMStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	result := gs.db.Find(&contacts)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve contacts: %w", result.Error)
	}
	if len(contacts) == 0 {
		return nil, fmt.Errorf("contact list is empty")
	}
	return contacts, nil
}

func (gs *GORMStore) GetById(id int) (*Contact, error) {
	var contact Contact
	result := gs.db.First(&contact, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrContactNotFound(id)
		}
		return nil, fmt.Errorf("failed to retrieve contact: %w", result.Error)
	}
	return &contact, nil
}

func (gs *GORMStore) Update(id int, newName, newEmail string) error {
	var contact Contact
	result := gs.db.First(&contact, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrContactNotFound(id)
		}
		return fmt.Errorf("failed to retrieve contact: %w", result.Error)
	}

	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}

	saveResult := gs.db.Save(&contact)
	if saveResult.Error != nil {
		return fmt.Errorf("failed to update contact: %w", saveResult.Error)
	}
	return nil
}

func (gs *GORMStore) Delete(id int) error {
	result := gs.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete contact: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrContactNotFound(id)
	}
	return nil
}
