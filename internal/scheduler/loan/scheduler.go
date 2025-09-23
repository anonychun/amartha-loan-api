package loan

import (
	"context"
	"log"
	"sync"

	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func (s *Scheduler) SendInvestedNotification(ctx context.Context) error {
	loans, err := s.repository.Loan.FindAllByStatusAndIsInvestedNotificationSentFalse(ctx, entity.LoanStatusInvested)
	if err != nil {
		return err
	}

	loanIds := lo.Map(loans, func(loan *entity.Loan, _ int) uuid.UUID {
		return loan.Id
	})

	investments, err := s.repository.Investment.FindAllByLoanIdIn(ctx, loanIds)
	if err != nil {
		return err
	}

	investorIds := lo.Map(investments, func(investment *entity.Investment, _ int) uuid.UUID {
		return investment.InvestorId
	})

	investors, err := s.repository.Investor.FindAllByIds(ctx, investorIds)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, investor := range investors {
		wg.Go(func() {
			log.Printf("Sending invested notification to investor %s (%s)\n", investor.Name, investor.EmailAddress)
		})
	}

	wg.Wait()
	return s.repository.Loan.UpdateIsInvestedNotificationSentByIdIn(ctx, true, loanIds)
}
