package current

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
	"gorm.io/gorm"
)

type key int

const (
	txKey key = iota
	adminKey
	borrowerKey
	investorKey
)

func Tx(ctx context.Context) *gorm.DB {
	tx, _ := ctx.Value(txKey).(*gorm.DB)
	return tx
}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func Admin(ctx context.Context) *entity.Admin {
	admin, _ := ctx.Value(adminKey).(*entity.Admin)
	return admin
}

func SetAdmin(ctx context.Context, admin *entity.Admin) context.Context {
	return context.WithValue(ctx, adminKey, admin)
}

func Borrower(ctx context.Context) *entity.Borrower {
	borrower, _ := ctx.Value(borrowerKey).(*entity.Borrower)
	return borrower
}

func SetBorrower(ctx context.Context, borrower *entity.Borrower) context.Context {
	return context.WithValue(ctx, borrowerKey, borrower)
}

func Investor(ctx context.Context) *entity.Investor {
	investor, _ := ctx.Value(investorKey).(*entity.Investor)
	return investor
}

func SetInvestor(ctx context.Context, investor *entity.Investor) context.Context {
	return context.WithValue(ctx, investorKey, investor)
}
