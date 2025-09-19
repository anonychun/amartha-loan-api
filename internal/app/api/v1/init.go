package v1

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/borrower"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/investor"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1"
	HandlerInjectorName = "handler.api.v1"
)

type Usecase struct {
	Admin    *admin.Usecase
	Borrower *borrower.Usecase
	Investor *investor.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin:    do.MustInvokeNamed[*admin.Usecase](i, admin.UsecaseInjectorName),
		Borrower: do.MustInvokeNamed[*borrower.Usecase](i, borrower.UsecaseInjectorName),
		Investor: do.MustInvokeNamed[*investor.Usecase](i, investor.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Admin    *admin.Handler
	Borrower *borrower.Handler
	Investor *investor.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Admin:    do.MustInvokeNamed[*admin.Handler](i, admin.HandlerInjectorName),
		Borrower: do.MustInvokeNamed[*borrower.Handler](i, borrower.HandlerInjectorName),
		Investor: do.MustInvokeNamed[*investor.Handler](i, investor.HandlerInjectorName),
	}, nil
}
