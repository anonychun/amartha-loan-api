package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type BorrowerSession struct {
	Id         uuid.UUID
	BorrowerId uuid.UUID
	Borrower   *Borrower
	Token      string
	IpAddress  string
	UserAgent  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (bs *BorrowerSession) BeforeCreate(tx *gorm.DB) error {
	bs.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (bs *BorrowerSession) GenerateToken() {
	bs.Token = ulid.Make().String()
}
