package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"gorm.io/gorm"
)

type StringIDModel struct {
	ID        string `gorm:"primaryKey;size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *StringIDModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID, err = generateRandomID()
	}

	return nil
}

func generateRandomID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
