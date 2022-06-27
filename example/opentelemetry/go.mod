module github.com/uptrace/go-clickhouse/example/opentelemetry

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

replace github.com/uptrace/go-clickhouse/chotel => ../../chotel

exclude go.opentelemetry.io/proto/otlp v0.15.0

require (
	github.com/brianvoe/gofakeit/v5 v5.11.2
	github.com/uptrace/go-clickhouse v0.2.8
	github.com/uptrace/go-clickhouse/chotel v0.2.8
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.14
	go.opentelemetry.io/otel v1.7.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/uptrace/uptrace-go v1.7.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.32.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.7.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	go.opentelemetry.io/proto/otlp v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220624220833-87e55d714810 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220624142145-8cd45d7dbd1f // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
