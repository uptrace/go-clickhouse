package main

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
)

type Model struct {
	ch.CHModel `ch:"partition:toYYYYMM(time)"`

	ID   uint64
	Text string    `ch:",lc"`
	Time time.Time `ch:",pk"`
}

func main() {
	ctx := context.Background()

	db := ch.Connect(ch.WithDatabase("test"))
	db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	var num int
	if err := db.QueryRowContext(ctx, "SELECT 123").Scan(&num); err != nil {
		panic(err)
	}
	fmt.Println(num)

	if err := db.ResetModel(ctx, (*Model)(nil)); err != nil {
		panic(err)
	}

	src := &Model{ID: 1, Text: "hello", Time: time.Now()}
	if _, err := db.NewInsert().Model(src).Exec(ctx); err != nil {
		panic(err)
	}

	dest := new(Model)
	if err := db.NewSelect().Model(dest).Where("id = ?", src.ID).Limit(1).Scan(ctx); err != nil {
		panic(err)
	}
	fmt.Println(dest)
}
