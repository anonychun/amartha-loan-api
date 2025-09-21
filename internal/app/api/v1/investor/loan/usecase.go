package loan

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/samber/lo"
)

func (u *Usecase) FindAll(ctx context.Context) ([]*LoanDto, error) {
	loans, err := u.repository.Loan.FindAllByStatusInOrderByIdDesc(ctx, []entity.LoanStatus{
		entity.LoanStatusApproved,
		entity.LoanStatusInvested,
		entity.LoanStatusDisbursed,
	})
	if err != nil {
		return nil, err
	}

	res := lo.Map(loans, func(loan *entity.Loan, _ int) *LoanDto {
		return ToLoanDto(loan)
	})

	return res, nil
}

func (u *Usecase) Invest(ctx context.Context, req InvestRequest) (*LoanDto, error) {
	loan, err := u.repository.Loan.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrLoanNotFound
	} else if err != nil {
		return nil, err
	}

	if loan.Status != entity.LoanStatusApproved {
		return nil, consts.ErrLoanNotAvailableForInvestment
	}

	err = repository.Transaction(ctx, func(ctx context.Context) error {
		totalInvestmentAmount, err := u.repository.Investment.SumOfAmountsByLoanId(ctx, loan.Id)
		if err != nil {
			return err
		}

		if req.Amount > loan.PrincipalAmount-totalInvestmentAmount {
			return consts.ErrInvestmentAmountExceedsAvailableAmount
		}

		investment := &entity.Investment{
			LoanId:     loan.Id,
			InvestorId: current.Investor(ctx).Id,
			Amount:     req.Amount,
		}

		err = u.repository.Investment.Create(ctx, investment)
		if err != nil {
			return err
		}

		if investment.Amount+totalInvestmentAmount == loan.PrincipalAmount {
			loan.Status = entity.LoanStatusInvested

			err = u.repository.Loan.Update(ctx, loan)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ToLoanDto(loan), nil
}
