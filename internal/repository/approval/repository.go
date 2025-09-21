package approval

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) Create(ctx context.Context, approval *entity.Approval) error {
	return r.sql.DB(ctx).Create(approval).Error
}
