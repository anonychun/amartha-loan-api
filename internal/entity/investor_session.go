package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type InvestorSession struct {
	Id         uuid.UUID
	InvestorId uuid.UUID
	Investor   *Investor
	Token      string
	IpAddress  string
	UserAgent  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (is *InvestorSession) BeforeCreate(tx *gorm.DB) error {
	is.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (is *InvestorSession) GenerateToken() {
	is.Token = ulid.Make().String()
}
