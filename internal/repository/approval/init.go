package approval

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/db"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, RepositoryInjectorName, NewRepository)
}

const RepositoryInjectorName = "repository.approval"

type Repository struct {
	sql *db.Sql
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		sql: do.MustInvoke[*db.Sql](i),
	}, nil
}
