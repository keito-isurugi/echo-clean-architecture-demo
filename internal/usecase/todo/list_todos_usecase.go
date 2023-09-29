package usecase

import (
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
	"context"

	"go.uber.org/zap"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/env"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/logger"
)

type ListTodosUsecase interface {
	Exec(ctx context.Context) (entity.ListTodos, error)
}

type listTodosUsecaseImpl struct {
	todoRepo domain.TodoRepository
}

func NewListTodosUsecase(todoRepo domain.TodoRepository) ListTodosUsecase {
	return &listTodosUsecaseImpl{
		todoRepo: todoRepo,
	}
}

func (g *listTodosUsecaseImpl) Exec(ctx context.Context) (entity.ListTodos, error) {
	// 処理を書く

	ev, _ := env.NewValue()

	zapLogger, _ := logger.NewLogger(true)
	defer func() { _ = zapLogger.Sync() }()
	zapLogger.Info("Server Start!", zap.Any("env", ev))

	dbClient, dberr := db.NewClient(&ev.DB, zapLogger)
	if dberr != nil {
		return nil, dberr
	}

	var todos entity.ListTodos
	if err := dbClient.Conn(ctx).
	Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil


}