module github.com/uptrace/go-clickhouse/example/basic

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

require (
	github.com/uptrace/go-clickhouse v0.2.8
	github.com/uptrace/go-clickhouse/chdebug v0.2.8
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d // indirect
	golang.org/x/sys v0.0.0-20220624220833-87e55d714810 // indirect
)
