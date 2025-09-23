package repository

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/db"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	sql, err := do.Invoke[*db.Sql](bootstrap.Injector)
	if err != nil {
		return err
	}

	return sql.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = current.SetTx(ctx, tx)
		return fn(ctx)
	})
}
