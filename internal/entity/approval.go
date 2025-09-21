package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Approval struct {
	Id        uuid.UUID
	LoanId    uuid.UUID
	Loan      *Loan
	AdminId   uuid.UUID
	Admin     *Admin
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Approval) BeforeCreate(tx *gorm.DB) error {
	a.Id = uuid.Must(uuid.NewV7())
	return nil
}
