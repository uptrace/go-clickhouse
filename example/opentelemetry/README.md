# Example for go-clickhouse OpenTelemetry instrumentation

See
[Performance and errors monitoring](https://clickhouse.uptrace.dev/guide/clickhouse-monitoring-performance.html)
for details.

You can run this example with different OpenTelemetry exporters by providing environment variables.

**Stdout** exporter (default):

```shell
go run .
```

**Jaeger** exporter:

```shell
OTEL_EXPORTER_JAEGER_ENDPOINT=http://localhost:14268/api/traces go run .
```

[Uptrace](https://github.com/uptrace/uptrace) exporter:

```shell
UPTRACE_DSN="https://<token>@uptrace.dev/<project_id>" go run .
```

## Links

- [Find instrumentations](https://opentelemetry.uptrace.dev/instrumentations/?lang=go)
- [OpenTelemetry Go Tracing API](https://opentelemetry.uptrace.dev/guide/go-tracing.html)
