package postgres

import (
	"context"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
	"go.uber.org/zap"
)

type todoRepository struct {
	dbClient db.Client
	zapLogger *zap.Logger
}

func NewTodoRepository(dbClient db.Client, zapLogger *zap.Logger) domain.TodoRepository {
	return &todoRepository{
		dbClient: dbClient,
		zapLogger: zapLogger,
	}
}

func (t *todoRepository) ListTodos(ctx context.Context) (entity.ListTodos, error) {
	var todos entity.ListTodos
	if err := t.dbClient.Conn(ctx).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}