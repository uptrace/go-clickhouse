package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/go-clickhouse/ch"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *ch.DB) error {
		fmt.Print(" [up migration] ")
		return nil
	}, func(ctx context.Context, db *ch.DB) error {
		fmt.Print(" [down migration] ")
		return nil
	})
}
