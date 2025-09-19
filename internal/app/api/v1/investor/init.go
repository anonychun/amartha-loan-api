package investor

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/investor/auth"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1.investor"
	HandlerInjectorName = "handler.api.v1.investor"
)

type Usecase struct {
	Auth *auth.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Auth: do.MustInvokeNamed[*auth.Usecase](i, auth.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Auth *auth.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Auth: do.MustInvokeNamed[*auth.Handler](i, auth.HandlerInjectorName),
	}, nil
}
