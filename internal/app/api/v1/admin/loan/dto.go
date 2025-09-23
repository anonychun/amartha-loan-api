package loan

import (
	"mime/multipart"
	"time"

	"github.com/anonychun/amartha-loan-api/internal/dto"
	"github.com/anonychun/amartha-loan-api/internal/entity"
)

type LoanDto struct {
	Id              string          `json:"id"`
	BorrowerId      string          `json:"borrowerId"`
	PrincipalAmount int64           `json:"principalAmount"`
	Status          string          `json:"status"`
	AgreementLetter *dto.Attachment `json:"agreementLetter"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
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

type FindByIdRequest struct {
	Id string `param:"id"`
}

type ApproveRequest struct {
	Id           string                `param:"id"`
	ProofOfVisit *multipart.FileHeader `form:"proofOfVisit"`
}

type DisburseRequest struct {
	Id              string                `param:"id"`
	AgreementLetter *multipart.FileHeader `form:"agreementLetter"`
}
