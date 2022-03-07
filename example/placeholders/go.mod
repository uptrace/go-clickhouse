module github.com/uptrace/go-clickhouse/example/placeholders

go 1.18

replace github.com/uptrace/go-clickhouse/ch => ../../ch

replace github.com/uptrace/go-clickhouse/extra/chdebug => ../../extra/chdebug

require (
	github.com/uptrace/go-clickhouse/ch v0.0.0-20220307120528-f94409373c22
	github.com/uptrace/go-clickhouse/extra/chdebug v0.0.0-20220307120528-f94409373c22
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	golang.org/x/exp v0.0.0-20220307080910-a2e15db56b5f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)
