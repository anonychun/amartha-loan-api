package loan

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1.borrower.loan"
	HandlerInjectorName = "handler.api.v1.borrower.loan"
)

type Usecase struct {
	repository *repository.Repository
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		repository: do.MustInvoke[*repository.Repository](i),
	}, nil
}

type Handler struct {
	usecase *Usecase
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		usecase: do.MustInvokeNamed[*Usecase](i, UsecaseInjectorName),
	}, nil
}
