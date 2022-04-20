package trace

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"io"
	"log"
	"os"
)

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(

		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func consoleExporter() (trace.SpanExporter, error) {
	return stdouttrace.New(
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
func InitTrace(f *os.File) *trace.TracerProvider {

	exp, err := newExporter(f)
	if err != nil {
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(1)),
		trace.WithResource(newResource()),
	)
	tp.RegisterSpanProcessor(trace.NewSimpleSpanProcessor(exp))
	return tp
}

func Ttt(f *os.File) *trace.TracerProvider {

	exp, err := newExporter(f)
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(1)),
		trace.WithResource(newResource()),
	)
	tp.RegisterSpanProcessor(trace.NewSimpleSpanProcessor(exp))
	return tp
}

func InitTestTrace() *trace.TracerProvider {
	l := log.New(os.Stdout, "", 0)
	exp, err := consoleExporter()
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(1)),
		trace.WithResource(newResource()),
	)
	tp.RegisterSpanProcessor(trace.NewSimpleSpanProcessor(exp))

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	return tp
}
