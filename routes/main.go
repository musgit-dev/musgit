package routes

import (
	"context"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"

	"github.com/musgit-dev/musgit/internal/musgit"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run(
	ctx context.Context,
	mg *musgit.Musgit, tracer trace.Tracer) {

	getRoutes(ctx, mg, tracer)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

func getRoutes(
	ctx context.Context,
	mg *musgit.Musgit,
	tracer trace.Tracer,
) {
	ctx, span := tracer.Start(ctx, "get-routes")
	defer span.End()
	api := router.Group("/api")
	api.Use(otelgin.Middleware("musgit-service"))
	lessonRoutes(mg, api)
	pieceRoutes(mg, api)
}
