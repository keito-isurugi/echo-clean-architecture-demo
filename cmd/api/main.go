package main

import (
	"go.uber.org/zap"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/aws"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/env"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/logger"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/server"
)

//	@Summary		Swagger Example API
//	@version		v1
//	@description	BA Portal Replace API
//	@host			localhost

func main() {
	ev, _ := env.NewValue()

	zapLogger, _ := logger.NewLogger(true)
	defer func() { _ = zapLogger.Sync() }()
	zapLogger.Info("Server Start!", zap.Any("env", ev))

	dbClient, err := db.NewClient(&ev.DB, zapLogger)
	if err != nil {
		zapLogger.Error(err.Error())
	}

	awsClient, err := aws.NewAWSSession(ev)
	if err != nil {
		zapLogger.Error(err.Error())
	}

	router := server.SetupRouter(ev, dbClient, awsClient, zapLogger)

	router.Logger.Fatal(router.Start(":" + ev.ServerPort))
}
