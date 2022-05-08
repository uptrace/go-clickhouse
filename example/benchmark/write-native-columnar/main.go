package main

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
)

type Model struct {
	ch.CHModel `ch:",columnar,engine:Null()"`

	Col1 []uint64
	Col2 []string
	Col3 [][]uint8
	Col4 []time.Time
}

func benchmark(ctx context.Context, db *ch.DB) error {
	var model Model

	for i := 0; i < 1_000_000; i++ {
		model.Col1 = append(model.Col1, uint64(i))
		model.Col2 = append(model.Col2, "Golang SQL database driver")
		model.Col3 = append(model.Col3, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9})
		model.Col4 = append(model.Col4, time.Now())
	}

	_, err := db.NewInsert().Model(&model).Exec(ctx)
	return err
}

func main() {
	ctx := context.Background()

	db := ch.Connect(
		ch.WithDatabase("test"),
		ch.WithCompression(false),
	)

	if err := db.ResetModel(ctx, (*Model)(nil)); err != nil {
		panic(err)
	}

	start := time.Now()
	if err := benchmark(ctx, db); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))
}
