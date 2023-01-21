module github.com/uptrace/go-clickhouse/ch/internal/bench

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

require github.com/uptrace/go-clickhouse v0.3.0

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	golang.org/x/exp v0.0.0-20230118134722-a68e582fa157 // indirect
)
