package scheduler

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/scheduler/loan"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewScheduler)
}

type Scheduler struct {
	Loan *loan.Scheduler
}

func NewScheduler(i do.Injector) (*Scheduler, error) {
	return &Scheduler{
		Loan: do.MustInvoke[*loan.Scheduler](i),
	}, nil
}
