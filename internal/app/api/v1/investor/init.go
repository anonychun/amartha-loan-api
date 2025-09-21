package investor

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/investor/auth"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/investor/loan"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Auth *auth.Usecase
	Loan *loan.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Auth: do.MustInvoke[*auth.Usecase](i),
		Loan: do.MustInvoke[*loan.Usecase](i),
	}, nil
}

type Handler struct {
	Auth *auth.Handler
	Loan *loan.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Auth: do.MustInvoke[*auth.Handler](i),
		Loan: do.MustInvoke[*loan.Handler](i),
	}, nil
}
