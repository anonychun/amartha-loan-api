package auth

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewMiddleware)
}

type Middleware struct {
	repository *repository.Repository
}

func NewMiddleware(i do.Injector) (*Middleware, error) {
	return &Middleware{
		repository: do.MustInvoke[*repository.Repository](i),
	}, nil
}
