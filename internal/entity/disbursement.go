package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Disbursement struct {
	Id        uuid.UUID
	LoanId    uuid.UUID
	Loan      *Loan
	AdminId   uuid.UUID
	Admin     *Admin
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *Disbursement) BeforeCreate(tx *gorm.DB) error {
	d.Id = uuid.Must(uuid.NewV7())
	return nil
}
