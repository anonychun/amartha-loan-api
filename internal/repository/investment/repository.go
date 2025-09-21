package investment

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/google/uuid"
)

func (r *Repository) Create(ctx context.Context, investment *entity.Investment) error {
	return r.sql.DB(ctx).Create(investment).Error
}

func (r *Repository) SumOfAmountsByLoanId(ctx context.Context, loanId uuid.UUID) (int64, error) {
	var total int64
	err := r.sql.DB(ctx).Model(&entity.Investment{}).Select("COALESCE(SUM(amount), 0)").Where("loan_id = ?", loanId).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}
