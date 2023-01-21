package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var dirFlag = flag.String("dir", "", "destination directory name")

var (
	//go:embed column.tpl
	columnSrc string
	//go:embed column_safe.tpl
	columnSafeSrc string
	//go:embed column_unsafe.tpl
	columnUnsafeSrc string
)

var columns []ColumnType

type ColumnType struct {
	Name      string
	GoType    string
	GoReflect string
	Size      int
	IsCustom  bool
}

func init() {
	for _, size := range []int{8, 16, 32, 64} {
		columns = append(columns, ColumnType{
			Name:      fmt.Sprintf("Int%d", size),
			GoType:    fmt.Sprintf("int%d", size),
			GoReflect: "Int",
			Size:      size,
		}, ColumnType{
			Name:      fmt.Sprintf("UInt%d", size),
			GoType:    fmt.Sprintf("uint%d", size),
			GoReflect: "Uint",
			Size:      size,
		})
	}
	for _, size := range []int{32, 64} {
		columns = append(columns, ColumnType{
			Name:      fmt.Sprintf("Float%d", size),
			GoType:    fmt.Sprintf("float%d", size),
			GoReflect: "Float",
			Size:      size,
		})
	}

	columns = append(columns, ColumnType{
		Name:      "Bool",
		GoType:    "bool",
		GoReflect: "Bool",
	}, ColumnType{
		Name:      "String",
		GoType:    "string",
		GoReflect: "String",
	}, ColumnType{
		Name:      "Bytes",
		GoType:    "[]byte",
		GoReflect: "Bytes",
	}, ColumnType{
		Name:      "Enum",
		GoType:    "string",
		GoReflect: "String",
		IsCustom:  true,
	}, ColumnType{
		Name:   "DateTime",
		GoType: "time.Time",
	})
}

func main() {
	flag.Parse()

	for _, v := range []struct {
		name string
		tpl  string
	}{
		{"column_gen", columnSrc},
		{"column_safe_gen", columnSafeSrc},
		{"column_unsafe_gen", columnUnsafeSrc},
	} {
		dest := filepath.Join(*dirFlag, v.name+".go")
		tpl := template.Must(template.New("").Parse(v.tpl))
		if err := gen(dest, tpl, columns); err != nil {
			log.Fatalf("%s: %s", v.name, err)
		}
	}
}

func gen(filePath string, tpl *template.Template, vars any) error {
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, vars); err != nil {
		return err
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		data = buf.Bytes()
	}

	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return err
	}

	return nil
}
