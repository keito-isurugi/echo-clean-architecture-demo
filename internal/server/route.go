package server

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/handler"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/env"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/postgres"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/server/middleware"
)

func SetupRouter(ev *env.Values, dbClient db.Client, awsClient s3iface.S3API, zapLogger *zap.Logger) *echo.Echo {
	e := echo.New()
	// middleware
	e.Use(middleware.NewLogging(zapLogger))
	e.Use(middleware.NewTracing())
	// e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	//postgres
	menuMasterRepo := postgres.NewMenuMasterRepository(dbClient, zapLogger)
	todoRepo := postgres.NewTodoRepository(dbClient, zapLogger)

	//handler
	menuMasterHandler := handler.NewMenuMasterHandler(menuMasterRepo, zapLogger)
	todoHandler := handler.NewTodoHandler(todoRepo, zapLogger)

	// menu_masters
	menuMasterGroup := e.Group("/menu_masters")
	menuMasterGroup.GET("", menuMasterHandler.ListMenuMasters)
	menuMasterGroup.POST("", menuMasterHandler.RegisterMenuMaster)

	// // hoge
	hogeGroup := e.Group("/hoge")
	hogeGroup.GET("", handler.Hoge)

	// todos
	todoGroup := e.Group("/todos")
	todoGroup.GET("", todoHandler.ListTodos)

	for _, route := range e.Routes() {
		zapLogger.Info("route", zap.String("method", route.Method), zap.String("path", route.Path))
	}

	return e
}
