package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Investment struct {
	Id         uuid.UUID
	LoanId     uuid.UUID
	Loan       *Loan
	InvestorId uuid.UUID
	Investor   *Investor
	Amount     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (i *Investment) BeforeCreate(tx *gorm.DB) error {
	i.Id = uuid.Must(uuid.NewV7())
	return nil
}
