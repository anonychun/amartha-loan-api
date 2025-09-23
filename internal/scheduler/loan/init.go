package loan

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewScheduler)
}

type Scheduler struct {
	repository *repository.Repository
}

func NewScheduler(i do.Injector) (*Scheduler, error) {
	return &Scheduler{
		repository: do.MustInvoke[*repository.Repository](i),
	}, nil
}
