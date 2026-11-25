package cmd

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/musgit-dev/musgit/internal/musgit"
	"github.com/musgit-dev/musgit/routes"
	"github.com/musgit-dev/musgit/tracing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tracer trace.Tracer

func startServer(dbUri string) {
	ctx := context.Background()

	tp := tracing.NewTracerProvider(ctx)

	defer func() { _ = tp.Shutdown(ctx) }()
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer("github.com/musgit-dev/musgit")
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	if dbUri == "" {
		log.Fatal("DB not provided.")
	}
	mg := musgit.New(dbUri)

	routes.Run(ctx, mg, tracer)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts server",
	Run: func(cmd *cobra.Command, args []string) {
		dbUri := viper.GetString("db-uri")
		startServer(dbUri)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
