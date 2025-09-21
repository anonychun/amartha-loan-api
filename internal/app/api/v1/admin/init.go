package admin

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin/admin"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin/auth"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin/loan"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1.admin"
	HandlerInjectorName = "handler.api.v1.admin"
)

type Usecase struct {
	Admin *admin.Usecase
	Auth  *auth.Usecase
	Loan  *loan.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin: do.MustInvokeNamed[*admin.Usecase](i, admin.UsecaseInjectorName),
		Auth:  do.MustInvokeNamed[*auth.Usecase](i, auth.UsecaseInjectorName),
		Loan:  do.MustInvokeNamed[*loan.Usecase](i, loan.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Auth  *auth.Handler
	Admin *admin.Handler
	Loan  *loan.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Admin: do.MustInvokeNamed[*admin.Handler](i, admin.HandlerInjectorName),
		Auth:  do.MustInvokeNamed[*auth.Handler](i, auth.HandlerInjectorName),
		Loan:  do.MustInvokeNamed[*loan.Handler](i, loan.HandlerInjectorName),
	}, nil
}
