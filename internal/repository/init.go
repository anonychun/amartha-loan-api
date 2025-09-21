package repository

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/db"
	"github.com/anonychun/amartha-loan-api/internal/repository/admin"
	"github.com/anonychun/amartha-loan-api/internal/repository/admin_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/approval"
	"github.com/anonychun/amartha-loan-api/internal/repository/borrower"
	"github.com/anonychun/amartha-loan-api/internal/repository/borrower_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/disbursement"
	"github.com/anonychun/amartha-loan-api/internal/repository/investment"
	"github.com/anonychun/amartha-loan-api/internal/repository/investor"
	"github.com/anonychun/amartha-loan-api/internal/repository/investor_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/loan"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	Admin           *admin.Repository
	AdminSession    *admin_session.Repository
	Approval        *approval.Repository
	Borrower        *borrower.Repository
	BorrowerSession *borrower_session.Repository
	Disbursement    *disbursement.Repository
	Investment      *investment.Repository
	Investor        *investor.Repository
	InvestorSession *investor_session.Repository
	Loan            *loan.Repository
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		Admin:           do.MustInvokeNamed[*admin.Repository](i, admin.RepositoryInjectorName),
		AdminSession:    do.MustInvokeNamed[*admin_session.Repository](i, admin_session.RepositoryInjectorName),
		Approval:        do.MustInvokeNamed[*approval.Repository](i, approval.RepositoryInjectorName),
		Borrower:        do.MustInvokeNamed[*borrower.Repository](i, borrower.RepositoryInjectorName),
		BorrowerSession: do.MustInvokeNamed[*borrower_session.Repository](i, borrower_session.RepositoryInjectorName),
		Disbursement:    do.MustInvokeNamed[*disbursement.Repository](i, disbursement.RepositoryInjectorName),
		Investment:      do.MustInvokeNamed[*investment.Repository](i, investment.RepositoryInjectorName),
		Investor:        do.MustInvokeNamed[*investor.Repository](i, investor.RepositoryInjectorName),
		InvestorSession: do.MustInvokeNamed[*investor_session.Repository](i, investor_session.RepositoryInjectorName),
		Loan:            do.MustInvokeNamed[*loan.Repository](i, loan.RepositoryInjectorName),
	}, nil
}

func Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	sql, err := do.Invoke[*db.Sql](bootstrap.Injector)
	if err != nil {
		return err
	}

	return sql.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = current.SetTx(ctx, tx)
		return fn(ctx)
	})
}
