# Migrations example

To run this example, you need a ClickHouse database:

```shell
clickhouse-client -q "CREATE DATABASE test"
```

To run migrations:

```shell
go run . db migrate
```

To rollback migrations:

```shell
go run . db rollback
```

To view status of migrations:

```shell
go run . db status
```

To create a Go migration:

```shell
go run . db create_go go_migration_name
```

To create a SQL migration:

```shell
go run . db create_sql sql_migration_name
```

See [ClickHouse migrations](https://clickhouse.uptrace.dev/guide/clickhouse-migrations.html) for
details.
