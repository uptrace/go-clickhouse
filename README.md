# ClickHouse client for Go 1.18+

[![build workflow](https://github.com/uptrace/go-clickhouse/actions/workflows/build.yml/badge.svg)](https://github.com/uptrace/go-clickhouse/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/go-clickhouse/ch)](https://pkg.go.dev/github.com/go-clickhouse/ch)
[![Documentation](https://img.shields.io/badge/ch-documentation-informational)](https://clickhouse.uptrace.dev/)

This client uses native protocol to communicate with ClickHouse server and requires Go 1.18+ in
order to use generics. This is not a database/sql driver, but the API is compatible.

Main features are:

- ClickHouse native protocol support and efficient column-oriented design.
- API compatible with database/sql.
- [Bun](https://github.com/uptrace/bun/)-like query builder.
- [Selecting](https://clickhouse.uptrace.dev/guide/query-select.html) into scalars, structs, maps,
  slices of maps/structs/scalars.
- `Array(T)` including nested arrays.
- Enums and `LowCardinality(String)`.
- `Nullable(T)` except `Nullable(Array(T))`.
- [Migrations](https://clickhouse.uptrace.dev/guide/migrations.html).
- [OpenTelemetry](https://clickhouse.uptrace.dev/guide/monitoring.html) support.
- In production at [Uptrace](https://uptrace.dev/)

Resources:

- [**Get started**](https://clickhouse.uptrace.dev/guide/getting-started.html)
- [Examples](https://github.com/uptrace/go-clickhouse/tree/master/example)
- [Discussions](https://github.com/uptrace/go-clickhouse/discussions)
- [Reference](https://pkg.go.dev/github.com/uptrace/go-clickhouse/ch)
- [Example app](https://github.com/uptrace/uptrace)

## Example

A [basic](example/basic) example:

```go
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
```
