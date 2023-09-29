package handler

import (
	"net/http"
	// "time"
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	// "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
	// usecase "github.com/kokopelli-inc/echo-ddd-demo/internal/usecase/menu_master"
	// "github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
	// "github.com/kokopelli-inc/echo-ddd-demo/internal/infra/env"
	// "github.com/kokopelli-inc/echo-ddd-demo/internal/infra/logger"

	usecase "github.com/kokopelli-inc/echo-ddd-demo/internal/usecase/todo"
)

type TodoHandler interface {
	ListTodos(echo.Context) error
}

type todoResponse struct {
	ID   int    `json:"id" example:"1"`
	Title string `json:"title" example:"Todoのタイトル"`
	Content string `json:"content" example:"Todoの内容"`
}

type listTodoResponse []todoResponse

type todoHandler struct {
	todoRepo domain.TodoRepository
	zapLogger *zap.Logger
}

func NewTodoHandler(
	todoRepo domain.TodoRepository, 
	zapLogger *zap.Logger,
) TodoHandler {
	return &todoHandler{
		todoRepo: todoRepo,
		zapLogger: zapLogger,
	}
}

func (h *todoHandler) ListTodos(c echo.Context) error {
	var ctx context.Context
	var todoRepo domain.TodoRepository
	uc := usecase.NewListTodosUsecase(todoRepo)
	res, err := uc.Exec(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
