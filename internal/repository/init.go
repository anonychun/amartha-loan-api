package repository

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/db"
	"github.com/anonychun/amartha-loan-api/internal/repository/admin"
	"github.com/anonychun/amartha-loan-api/internal/repository/admin_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/borrower"
	"github.com/anonychun/amartha-loan-api/internal/repository/borrower_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/investor"
	"github.com/anonychun/amartha-loan-api/internal/repository/investor_session"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	sql *db.Sql

	Admin           *admin.Repository
	AdminSession    *admin_session.Repository
	Borrower        *borrower.Repository
	BorrowerSession *borrower_session.Repository
	Investor        *investor.Repository
	InvestorSession *investor_session.Repository
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		sql: do.MustInvoke[*db.Sql](i),

		Admin:           do.MustInvokeNamed[*admin.Repository](i, admin.RepositoryInjectorName),
		AdminSession:    do.MustInvokeNamed[*admin_session.Repository](i, admin_session.RepositoryInjectorName),
		Borrower:        do.MustInvokeNamed[*borrower.Repository](i, borrower.RepositoryInjectorName),
		BorrowerSession: do.MustInvokeNamed[*borrower_session.Repository](i, borrower_session.RepositoryInjectorName),
		Investor:        do.MustInvokeNamed[*investor.Repository](i, investor.RepositoryInjectorName),
		InvestorSession: do.MustInvokeNamed[*investor_session.Repository](i, investor_session.RepositoryInjectorName),
	}, nil
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.sql.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = current.SetTx(ctx, tx)
		return fn(ctx)
	})
}
