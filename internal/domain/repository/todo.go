//go:generate mockgen -source=todo.go -destination=./mock/todo_mock.go
package domain

import (
	"context"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
)

type TodoRepository interface {
	ListTodos(ctx context.Context) (entity.ListTodos, error)
}
