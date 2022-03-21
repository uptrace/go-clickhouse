module github.com/uptrace/go-clickhouse/example/opentelemetry

go 1.18

replace github.com/uptrace/go-clickhouse => ../..

replace github.com/uptrace/go-clickhouse/chdebug => ../../chdebug

replace github.com/uptrace/go-clickhouse/chotel => ../../chotel

require (
	github.com/brianvoe/gofakeit/v5 v5.11.2
	github.com/uptrace/go-clickhouse v0.2.2
	github.com/uptrace/go-clickhouse/chotel v0.2.2
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.10
	go.opentelemetry.io/otel v1.5.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/uptrace/uptrace-go v1.4.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.5.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.5.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.5.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.5.0 // indirect
	go.opentelemetry.io/proto/otlp v0.12.0 // indirect
	golang.org/x/exp v0.0.0-20220318154914-8dddf5d87bd8 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
