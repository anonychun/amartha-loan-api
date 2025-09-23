package loan

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) FindById(ctx context.Context, id string) (*entity.Loan, error) {
	loan := &entity.Loan{}
	err := r.sql.DB(ctx).First(loan, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (r *Repository) FindByIdWithAgreementLetter(ctx context.Context, id string) (*entity.Loan, error) {
	loan := &entity.Loan{}
	err := r.sql.DB(ctx).Joins("AgreementLetter").First(loan, "loans.id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (r *Repository) FindAllOrderByIdDesc(ctx context.Context) ([]*entity.Loan, error) {
	loans := make([]*entity.Loan, 0)
	err := r.sql.DB(ctx).Order("id DESC").Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *Repository) FindAllByBorrowerIdOrderByIdDesc(ctx context.Context, borrowerId string) ([]*entity.Loan, error) {
	loans := make([]*entity.Loan, 0)
	err := r.sql.DB(ctx).Where("borrower_id = ?", borrowerId).Order("id DESC").Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *Repository) FindAllByStatusOrderByIdDesc(ctx context.Context, status entity.LoanStatus) ([]*entity.Loan, error) {
	loans := make([]*entity.Loan, 0)
	err := r.sql.DB(ctx).Where("status = ?", status).Order("id DESC").Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *Repository) FindAllByStatusInOrderByIdDesc(ctx context.Context, statuses []entity.LoanStatus) ([]*entity.Loan, error) {
	loans := make([]*entity.Loan, 0)
	err := r.sql.DB(ctx).Where("status IN ?", statuses).Order("id DESC").Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *Repository) Create(ctx context.Context, loan *entity.Loan) error {
	return r.sql.DB(ctx).Create(loan).Error
}

func (r *Repository) Update(ctx context.Context, loan *entity.Loan) error {
	return r.sql.DB(ctx).Save(loan).Error
}
