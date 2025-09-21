package loan

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/samber/lo"
)

func (u *Usecase) FindAll(ctx context.Context) ([]*LoanDto, error) {
	loans, err := u.repository.Loan.FindAllByBorrowerIdOrderByIdDesc(ctx, current.Borrower(ctx).Id.String())
	if err != nil {
		return nil, err
	}

	res := lo.Map(loans, func(loan *entity.Loan, _ int) *LoanDto {
		return ToLoanDto(loan)
	})

	return res, nil
}

func (u *Usecase) Create(ctx context.Context, req CreateRequest) (*LoanDto, error) {
	loan := &entity.Loan{
		BorrowerId:      current.Borrower(ctx).Id,
		PrincipalAmount: req.PrincipalAmount,
		Status:          entity.LoanStatusProposed,
	}

	err := u.repository.Loan.Create(ctx, loan)
	if err != nil {
		return nil, err
	}

	return ToLoanDto(loan), nil
}
