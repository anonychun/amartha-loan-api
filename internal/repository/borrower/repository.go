package borrower

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) FindAll(ctx context.Context) ([]*entity.Borrower, error) {
	borrowers := make([]*entity.Borrower, 0)
	err := r.sql.DB(ctx).Find(&borrowers).Error
	if err != nil {
		return nil, err
	}

	return borrowers, nil
}

func (r *Repository) FindById(ctx context.Context, id string) (*entity.Borrower, error) {
	borrower := &entity.Borrower{}
	err := r.sql.DB(ctx).First(borrower, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func (r *Repository) FindByEmailAddress(ctx context.Context, emailAddress string) (*entity.Borrower, error) {
	borrower := &entity.Borrower{}
	err := r.sql.DB(ctx).First(borrower, "email_address = ?", emailAddress).Error
	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func (r *Repository) Create(ctx context.Context, borrower *entity.Borrower) error {
	return r.sql.DB(ctx).Create(borrower).Error
}

func (r *Repository) Update(ctx context.Context, borrower *entity.Borrower) error {
	return r.sql.DB(ctx).Save(borrower).Error
}

func (r *Repository) ExistsById(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT 1 FROM borrowers WHERE id = ?", id).Scan(&exists).Error
	return exists, err
}

func (r *Repository) ExistsByEmailAddress(ctx context.Context, emailAddress string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT 1 FROM borrowers WHERE email_address = ?", emailAddress).Scan(&exists).Error
	return exists, err
}
