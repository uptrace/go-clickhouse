module github.com/uptrace/go-clickhouse/ch/internal/bench

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

require github.com/uptrace/go-clickhouse v0.2.7

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	golang.org/x/exp v0.0.0-20220516143420-24438e51023a // indirect
	golang.org/x/sys v0.0.0-20220517195934-5e4e11fc645e // indirect
)
