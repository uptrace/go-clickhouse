package main

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
)

const query = `
SELECT
	number
	, randomString(25)
	, array(1, 2, 3, 4, 5)
	, now()
FROM system.numbers LIMIT 1000000
`

func benchmark(ctx context.Context, db *ch.DB) error {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			col1 uint64
			col2 string
			col3 []uint8
			col4 time.Time
		)
		if err := rows.Scan(&col1, &col2, &col3, &col4); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ctx := context.Background()

	db := ch.Connect(
		ch.WithDatabase("test"),
		ch.WithCompression(false),
	)

	start := time.Now()
	if err := benchmark(ctx, db); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))
}
