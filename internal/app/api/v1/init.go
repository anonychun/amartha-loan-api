package v1

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/borrower"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/investor"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Admin    *admin.Usecase
	Borrower *borrower.Usecase
	Investor *investor.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin:    do.MustInvoke[*admin.Usecase](i),
		Borrower: do.MustInvoke[*borrower.Usecase](i),
		Investor: do.MustInvoke[*investor.Usecase](i),
	}, nil
}

type Handler struct {
	Admin    *admin.Handler
	Borrower *borrower.Handler
	Investor *investor.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Admin:    do.MustInvoke[*admin.Handler](i),
		Borrower: do.MustInvoke[*borrower.Handler](i),
		Investor: do.MustInvoke[*investor.Handler](i),
	}, nil
}
