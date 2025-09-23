package repository

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/repository/admin"
	"github.com/anonychun/amartha-loan-api/internal/repository/admin_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/approval"
	"github.com/anonychun/amartha-loan-api/internal/repository/attachment"
	"github.com/anonychun/amartha-loan-api/internal/repository/borrower"
	"github.com/anonychun/amartha-loan-api/internal/repository/borrower_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/disbursement"
	"github.com/anonychun/amartha-loan-api/internal/repository/investment"
	"github.com/anonychun/amartha-loan-api/internal/repository/investor"
	"github.com/anonychun/amartha-loan-api/internal/repository/investor_session"
	"github.com/anonychun/amartha-loan-api/internal/repository/loan"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	Admin           *admin.Repository
	AdminSession    *admin_session.Repository
	Approval        *approval.Repository
	Attachment      *attachment.Repository
	Borrower        *borrower.Repository
	BorrowerSession *borrower_session.Repository
	Disbursement    *disbursement.Repository
	Investment      *investment.Repository
	Investor        *investor.Repository
	InvestorSession *investor_session.Repository
	Loan            *loan.Repository
}

func NewRepository(i do.Injector) (*Repository, error) {
	return &Repository{
		Admin:           do.MustInvoke[*admin.Repository](i),
		AdminSession:    do.MustInvoke[*admin_session.Repository](i),
		Approval:        do.MustInvoke[*approval.Repository](i),
		Attachment:      do.MustInvoke[*attachment.Repository](i),
		Borrower:        do.MustInvoke[*borrower.Repository](i),
		BorrowerSession: do.MustInvoke[*borrower_session.Repository](i),
		Disbursement:    do.MustInvoke[*disbursement.Repository](i),
		Investment:      do.MustInvoke[*investment.Repository](i),
		Investor:        do.MustInvoke[*investor.Repository](i),
		InvestorSession: do.MustInvoke[*investor_session.Repository](i),
		Loan:            do.MustInvoke[*loan.Repository](i),
	}, nil
}
