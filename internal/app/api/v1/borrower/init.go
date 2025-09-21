package borrower

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/borrower/auth"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/borrower/loan"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1.borrower"
	HandlerInjectorName = "handler.api.v1.borrower"
)

type Usecase struct {
	Auth *auth.Usecase
	Loan *loan.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Auth: do.MustInvokeNamed[*auth.Usecase](i, auth.UsecaseInjectorName),
		Loan: do.MustInvokeNamed[*loan.Usecase](i, loan.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Auth *auth.Handler
	Loan *loan.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Auth: do.MustInvokeNamed[*auth.Handler](i, auth.HandlerInjectorName),
		Loan: do.MustInvokeNamed[*loan.Handler](i, loan.HandlerInjectorName),
	}, nil
}
