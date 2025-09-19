package borrower_session

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) FindByToken(ctx context.Context, token string) (*entity.BorrowerSession, error) {
	borrowerSession := &entity.BorrowerSession{}
	err := r.sql.DB(ctx).First(borrowerSession, "token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return borrowerSession, nil
}

func (r *Repository) Create(ctx context.Context, borrowerSession *entity.BorrowerSession) error {
	return r.sql.DB(ctx).Create(borrowerSession).Error
}

func (r *Repository) DeleteById(ctx context.Context, id string) error {
	return r.sql.DB(ctx).Delete(&entity.BorrowerSession{}, "id = ?", id).Error
}
