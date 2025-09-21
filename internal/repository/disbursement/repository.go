package disbursement

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) Create(ctx context.Context, disbursement *entity.Disbursement) error {
	return r.sql.DB(ctx).Create(disbursement).Error
}
