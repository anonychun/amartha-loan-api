package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Loan struct {
	Id              uuid.UUID
	BorrowerId      uuid.UUID
	Borrower        *Borrower
	PrincipalAmount int64
	Status          LoanStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type LoanStatus string

const (
	LoanStatusProposed  LoanStatus = "proposed"
	LoanStatusApproved  LoanStatus = "approved"
	LoanStatusInvested  LoanStatus = "invested"
	LoanStatusDisbursed LoanStatus = "disbursed"
)

func (l *Loan) BeforeCreate(tx *gorm.DB) error {
	l.Id = uuid.Must(uuid.NewV7())
	return nil
}
