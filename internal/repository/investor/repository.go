package investor

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/google/uuid"
)

func (r *Repository) FindAll(ctx context.Context) ([]*entity.Investor, error) {
	investors := make([]*entity.Investor, 0)
	err := r.sql.DB(ctx).Find(&investors).Error
	if err != nil {
		return nil, err
	}

	return investors, nil
}

func (r *Repository) FindById(ctx context.Context, id string) (*entity.Investor, error) {
	investor := &entity.Investor{}
	err := r.sql.DB(ctx).First(investor, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return investor, nil
}

func (r *Repository) FindByEmailAddress(ctx context.Context, emailAddress string) (*entity.Investor, error) {
	investor := &entity.Investor{}
	err := r.sql.DB(ctx).First(investor, "email_address = ?", emailAddress).Error
	if err != nil {
		return nil, err
	}

	return investor, nil
}

func (r *Repository) FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*entity.Investor, error) {
	investors := make([]*entity.Investor, 0)
	err := r.sql.DB(ctx).Where("id IN ?", ids).Find(&investors).Error
	if err != nil {
		return nil, err
	}

	return investors, nil
}

func (r *Repository) Create(ctx context.Context, investor *entity.Investor) error {
	return r.sql.DB(ctx).Create(investor).Error
}

func (r *Repository) Update(ctx context.Context, investor *entity.Investor) error {
	return r.sql.DB(ctx).Save(investor).Error
}

func (r *Repository) ExistsById(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT 1 FROM investors WHERE id = ?", id).Scan(&exists).Error
	return exists, err
}

func (r *Repository) ExistsByEmailAddress(ctx context.Context, emailAddress string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT 1 FROM investors WHERE email_address = ?", emailAddress).Scan(&exists).Error
	return exists, err
}
