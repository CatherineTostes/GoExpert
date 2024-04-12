package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/devfullcycle/desafio-opentelemetry/internal/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	resource2 "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	CreateServer()
}

func CreateServer() {
	initProvider()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Post("/", web.GetCep)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

func initProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	exporter, err := zipkin.New(os.Getenv("ZIPKIN_ENDPOINT"))
	if err != nil {
		log.Fatal("Fail to create zipkin exporter: %v", err)
	}

	res, err := resource2.New(ctx, resource2.WithAttributes(
		semconv.ServiceName("OTEL_SERVICE_NAME")))

	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}
