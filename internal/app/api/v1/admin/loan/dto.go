package loan

import (
	"time"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

type LoanDto struct {
	Id              string    `json:"id"`
	BorrowerId      string    `json:"borrowerId"`
	PrincipalAmount int64     `json:"principalAmount"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func ToLoanDto(loan *entity.Loan) *LoanDto {
	return &LoanDto{
		Id:              loan.Id.String(),
		BorrowerId:      loan.BorrowerId.String(),
		PrincipalAmount: loan.PrincipalAmount,
		Status:          string(loan.Status),
		CreatedAt:       loan.CreatedAt,
		UpdatedAt:       loan.UpdatedAt,
	}
}

type ApproveRequest struct {
	Id string `param:"id"`
}

type DisburseRequest struct {
	Id string `param:"id"`
}
