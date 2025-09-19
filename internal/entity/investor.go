package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Investor struct {
	Id             uuid.UUID
	Name           string
	EmailAddress   string
	PasswordDigest string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (i *Investor) BeforeCreate(tx *gorm.DB) error {
	i.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (i *Investor) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	i.PasswordDigest = string(hash)

	return nil
}

func (i *Investor) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(i.PasswordDigest), []byte(password))
}
