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
	//go:embed column_safe.tpl
	columnSafeSrc string
	//go:embed column_unsafe.tpl
	columnUnsafeSrc string
)

var types []ColumnType

type ColumnType struct {
	CHType string
	GoType string
	Size   int
}

func init() {
	for _, size := range []int{8, 16, 32, 64} {
		types = append(types, ColumnType{
			CHType: fmt.Sprintf("Int%d", size),
			GoType: fmt.Sprintf("int%d", size),
			Size:   size,
		}, ColumnType{
			CHType: fmt.Sprintf("UInt%d", size),
			GoType: fmt.Sprintf("uint%d", size),
			Size:   size,
		})
	}
	for _, size := range []int{32, 64} {
		types = append(types, ColumnType{
			Size:   size,
			CHType: fmt.Sprintf("Float%d", size),
			GoType: fmt.Sprintf("float%d", size),
		})
	}
}

func main() {
	flag.Parse()

	for _, v := range []struct {
		name string
		tpl  *template.Template
	}{
		{"column_safe_gen", template.Must(template.New("").Parse(columnSafeSrc))},
		{"column_unsafe_gen", template.Must(template.New("").Parse(columnUnsafeSrc))},
	} {
		if err := write(filepath.Join(*dirFlag, v.name+".go"), v.tpl, types); err != nil {
			log.Fatal(err)
		}
	}
}

func write(filePath string, tpl *template.Template, vars interface{}) error {
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, vars); err != nil {
		return err
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return err
	}

	return nil
}
