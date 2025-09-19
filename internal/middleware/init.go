package middleware

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/middleware/auth"
	"github.com/samber/do"
)

func init() {
	do.Provide(bootstrap.Injector, NewMiddleware)
}

type Middleware struct {
	Auth *auth.Middleware
}

func NewMiddleware(i *do.Injector) (*Middleware, error) {
	return &Middleware{
		Auth: do.MustInvokeNamed[*auth.Middleware](i, auth.MiddlewareInjectorName),
	}, nil
}
