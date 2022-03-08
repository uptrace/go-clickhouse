module github.com/uptrace/go-clickhouse/example/placeholders

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

require (
	github.com/uptrace/go-clickhouse/ch v0.0.0-20220308124651-82a5d8d72ef7
	github.com/uptrace/go-clickhouse/chdebug v0.0.0-20220308142411-e669ca4a7420
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	golang.org/x/exp v0.0.0-20220307200941-a1099baf94bf // indirect
	golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7 // indirect
)
