module github.com/uptrace/go-clickhouse/extra/chotel

go 1.18

replace github.com/uptrace/go-clickhouse/ch => ../../ch

replace github.com/uptrace/go-clickhouse/extra/chdebug => ../../extra/chdebug

require (
	github.com/uptrace/go-clickhouse/ch v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/trace v1.3.0
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/go-logr/logr v1.2.1 // indirect
	github.com/go-logr/stdr v1.2.0 // indirect
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.11 // indirect
	golang.org/x/exp v0.0.0-20211210185655-e05463a05a18 // indirect
)
