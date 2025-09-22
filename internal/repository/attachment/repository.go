package attachment

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (r *Repository) Create(ctx context.Context, attachment *entity.Attachment) error {
	return r.sql.DB(ctx).Create(attachment).Error
}
