# ClickHouse client for Go 1.18+

[![build workflow](https://github.com/uptrace/go-clickhouse/actions/workflows/build.yml/badge.svg)](https://github.com/uptrace/go-clickhouse/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/go-clickhouse/ch)](https://pkg.go.dev/github.com/go-clickhouse/ch)
[![Documentation](https://img.shields.io/badge/ch-documentation-informational)](https://clickhouse.uptrace.dev/)

This client uses native protocol to communicate with ClickHouse server. It requires Go 1.18+ in
order to use generics. This is not a database/sql driver, but the API is similar.

Main features are:

- Native protocol support.
- `database/sql`-like API.
- [Bun](https://github.com/uptrace/bun/)-like query builder.
- [Selecting](/example/basic/) into scalars, structs, maps, slices of maps/structs/scalars.
- Efficient inserts.
- `Array(*)` including nested arrays.
- Enums and `LowCardinality(String)`.
- Migrations.

Not supported:

- Nullable types.

Resources:

- [**Get started**](https://clickhouse.uptrace.dev/guide/getting-started.html)
- [Examples](https://github.com/uptrace/go-clickhouse/tree/master/example)
- [Discussions](https://github.com/uptrace/go-clickhouse/discussions)
- [Reference](https://pkg.go.dev/github.com/uptrace/go-clickhouse/ch)

## Installation

```go
go get github.com/uptrace/go-clickhouse/ch
```
