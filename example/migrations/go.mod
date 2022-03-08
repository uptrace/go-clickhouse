module github.com/uptrace/go-clickhouse/example/migrations

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

require (
	github.com/uptrace/go-clickhouse v0.0.0-20220308124651-82a5d8d72ef7
	github.com/uptrace/go-clickhouse/extra/chdebug v0.0.0-20220308120244-50192cd8fb1b
	github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20220307200941-a1099baf94bf // indirect
	golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7 // indirect
)