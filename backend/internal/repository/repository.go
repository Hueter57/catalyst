package repository

import (
	"context"

	"github.com/hueter57/catalyst/backend/internal/domain"
	"github.com/hueter57/catalyst/backend/internal/ent"
)

type Repository struct {
	c *ent.Client
}

func NewRepository(client *ent.Client) *Repository {
	return &Repository{
		c: client,
	}
}

// `atlas migrate diff` へのエイリアスです.
//
// `MigrationsDir` のディレクトリを参照してdiffの計算が行われます.
func (r *Repository) MigrateDiff(ctx context.Context, options ...domain.MigrateOption) error {
	return domain.MigrateDiff(ctx, r.c, options...)
}

// `atlas migrate apply` へのエイリアスです.
//
// `MigrationsDir` のディレクトリを参照してdiffの計算が行われます.
func (r *Repository) MigrateApply(ctx context.Context, options ...domain.MigrateOption) error {
	return domain.MigrateApply(ctx, r.c, options...)
}
