package worker

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/scheduler"
	"github.com/go-co-op/gocron/v2"
	"github.com/samber/do/v2"
)

func Start(ctx context.Context) error {
	s := do.MustInvoke[*scheduler.Scheduler](bootstrap.Injector)

	cron, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = cron.NewJob(
		gocron.CronJob("*/5 * * * *", false),
		gocron.NewTask(func() {
			s.Loan.SendInvestedNotification(ctx)
		}),
	)
	if err != nil {
		return err
	}

	cron.Start()
	<-ctx.Done()
	return cron.Shutdown()
}
