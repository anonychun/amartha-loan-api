package admin

import (
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin/admin"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin/auth"
	"github.com/anonychun/amartha-loan-api/internal/app/api/v1/admin/loan"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Admin *admin.Usecase
	Auth  *auth.Usecase
	Loan  *loan.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin: do.MustInvoke[*admin.Usecase](i),
		Auth:  do.MustInvoke[*auth.Usecase](i),
		Loan:  do.MustInvoke[*loan.Usecase](i),
	}, nil
}

type Handler struct {
	Auth  *auth.Handler
	Admin *admin.Handler
	Loan  *loan.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Admin: do.MustInvoke[*admin.Handler](i),
		Auth:  do.MustInvoke[*auth.Handler](i),
		Loan:  do.MustInvoke[*loan.Handler](i),
	}, nil
}
