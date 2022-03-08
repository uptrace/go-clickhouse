module github.com/uptrace/go-clickhouse/chotel

go 1.18

replace github.com/uptrace/go-clickhouse/ch => ../../ch

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

require (
	github.com/uptrace/go-clickhouse/ch v0.0.0-20220307120720-c1e94c04a777
	go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/trace v1.4.1
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	golang.org/x/exp v0.0.0-20220307200941-a1099baf94bf // indirect
)
