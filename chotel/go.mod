module github.com/uptrace/go-clickhouse/chotel

go 1.18

replace github.com/uptrace/go-clickhouse => ./..

replace github.com/uptrace/go-clickhouse/chdebug => ../chdebug

exclude go.opentelemetry.io/proto/otlp v0.15.0

require (
	github.com/uptrace/go-clickhouse v0.2.8
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/trace v1.8.0
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
)
