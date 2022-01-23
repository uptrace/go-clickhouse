package main

import (
	"context"
	"fmt"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
)

type User struct {
	ID     int64
	Name   string
	Emails []string
}

func main() {
	ctx := context.Background()

	db := ch.Connect(ch.WithDatabase("test"))
	db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))

	var tableName, tableAlias, pks, tablePKs, columns, tableColumns string

	if err := db.NewSelect().Model((*User)(nil)).
		ColumnExpr("'?TableName'").
		ColumnExpr("'?TableAlias'").
		ColumnExpr("'?PKs'").
		ColumnExpr("'?TablePKs'").
		ColumnExpr("'?Columns'").
		ColumnExpr("'?TableColumns'").
		ModelTableExpr("").
		Scan(ctx, &tableName, &tableAlias, &pks, &tablePKs, &columns, &tableColumns); err != nil {
		panic(err)
	}

	fmt.Println("tableName", tableName)
	fmt.Println("tableAlias", tableAlias)
	fmt.Println("pks", pks)
	fmt.Println("tablePKs", tablePKs)
	fmt.Println("columns", columns)
	fmt.Println("tableColumns", tableColumns)
}
