package investor_session

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) FindByToken(ctx context.Context, token string) (*entity.InvestorSession, error) {
	investorSession := &entity.InvestorSession{}
	err := r.sql.DB(ctx).First(investorSession, "token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return investorSession, nil
}

func (r *Repository) Create(ctx context.Context, investorSession *entity.InvestorSession) error {
	return r.sql.DB(ctx).Create(investorSession).Error
}

func (r *Repository) DeleteById(ctx context.Context, id string) error {
	return r.sql.DB(ctx).Delete(&entity.InvestorSession{}, "id = ?", id).Error
}
