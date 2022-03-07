module github.com/uptrace/go-clickhouse/ch

go 1.18

replace github.com/uptrace/go-clickhouse/extra/chdebug => ../extra/chdebug

require (
	github.com/bradleyjkemp/cupaloy v2.3.0+incompatible
	github.com/codemodus/kace v0.5.1
	github.com/go-pg/zerochecker v0.2.0
	github.com/jinzhu/inflection v1.0.0
	github.com/pierrec/lz4/v4 v4.1.14
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/go-clickhouse/extra/chdebug v0.0.0-00010101000000-000000000000
	golang.org/x/exp v0.0.0-20220307080910-a2e15db56b5f
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
