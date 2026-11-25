package tracing

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.43.0"
)

func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	spanExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(os.Getenv("OTLP_ENDPOINT")),
	)
	if err != nil {
		panic(err)
	}
	return spanExporter, err

}
func NewTracerProvider(ctx context.Context) *sdktrace.TracerProvider {
	exp, err := newExporter(ctx)
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("MusgitService"),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}
