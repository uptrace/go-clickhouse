package ch_test

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/ch/chschema"
)

func TestQuery(t *testing.T) {
	type Model struct {
		ID     uint64
		String string
		Bytes  []byte
	}

	queries := []func(db *ch.DB) chschema.QueryAppender{
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewCreateTable().Model((*Model)(nil))
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewDropTable().Model((*Model)(nil))
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewSelect().Model((*Model)(nil))
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewSelect().Model((*Model)(nil)).ExcludeColumn("bytes")
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewInsert().Model(new(Model))
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewTruncateTable().Model(new(Model))
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewSelect().
				Model((*Model)(nil)).
				Setting("max_rows_to_read = 100")
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewSelect().
				Model((*Model)(nil)).
				Setting("max_rows_to_read = 100").
				Setting("read_overflow_mode = 'break'")
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewInsert().
				TableExpr("dest").
				TableExpr("src").
				Where("_part = ?", "part_name").
				Setting("max_threads = 1").
				Setting("max_insert_threads = 1").
				Setting("max_execution_time = 0")
		},
		func(db *ch.DB) chschema.QueryAppender {
			return db.NewSelect().
				Model((*Model)(nil)).
				Sample("?", 1000)
		},
		func(db *ch.DB) chschema.QueryAppender {
			type Model struct {
				ch.CHModel `ch:"table:spans,partition:toYYYYMM(time)"`

				ID   uint64
				Text string    `ch:",lc"` // low cardinality column
				Time time.Time `ch:",pk"` // ClickHouse primary key for order by
			}
			return db.NewCreateTable().Model((*Model)(nil)).
				TTL("time + INTERVAL 30 DAY DELETE").
				Partition("toDate(time)").
				Setting("ttl_only_drop_parts = 1")
		},
	}

	db := chDB()
	defer db.Close()

	snapshotsDir := filepath.Join("testdata", "snapshots")
	snapshot := cupaloy.New(cupaloy.SnapshotSubdirectory(snapshotsDir))

	for i, fn := range queries {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			q := fn(db)

			query, err := q.AppendQuery(db.Formatter(), nil)
			if err != nil {
				snapshot.SnapshotT(t, err.Error())
			} else {
				snapshot.SnapshotT(t, string(query))
			}
		})
	}
}
