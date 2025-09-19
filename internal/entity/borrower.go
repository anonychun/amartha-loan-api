package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Borrower struct {
	Id             uuid.UUID
	Name           string
	EmailAddress   string
	PasswordDigest string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (b *Borrower) BeforeCreate(tx *gorm.DB) error {
	b.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (b *Borrower) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	b.PasswordDigest = string(hash)

	return nil
}

func (b *Borrower) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(b.PasswordDigest), []byte(password))
}
