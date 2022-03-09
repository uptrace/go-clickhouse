package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chotel"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var tracer = otel.Tracer("github.com/uptrace/go-clickhouse/example/opentelemetry")

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	db := ch.Connect(ch.WithDatabase("test"))
	db.AddQueryHook(chotel.NewQueryHook())
	// db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	if err := db.ResetModel(ctx, (*Model)(nil)); err != nil {
		panic(err)
	}

	ctx, span := tracer.Start(ctx, "handleRequest")
	defer span.End()

	if err := handleRequest(ctx, db); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	otelplay.PrintTraceID(ctx)
}

type Model struct {
	ch.CHModel `ch:"partition:toYYYYMM(time)"`

	ID   uint64
	Name string    `ch:",lc"`
	Time time.Time `ch:",pk"`
}

func handleRequest(ctx context.Context, db *ch.DB) error {
	model := &Model{
		ID:   rand.Uint64(),
		Name: gofakeit.Name(),
		Time: time.Now(),
	}
	if _, err := db.NewInsert().Model(model).Exec(ctx); err != nil {
		return err
	}

	// Check that data can be selected without any errors.
	if err := db.NewSelect().Model(model).Limit(1).Scan(ctx); err != nil {
		return err
	}

	return nil
}
