package usecase

import (
	"context"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
)

type ListMenuMastersUsecase interface {
	Exec(ctx context.Context) (entity.ListMenuMasters, error)
}

type listMenuMastersUsecaseImpl struct {
	menuMasterRepo domain.MenuMasterRepository
}

func NewListMenuMasterUsecase(menuMasterRepo domain.MenuMasterRepository) ListMenuMastersUsecase {
	return &listMenuMastersUsecaseImpl{
		menuMasterRepo: menuMasterRepo,
	}
}

func (g *listMenuMastersUsecaseImpl) Exec(ctx context.Context) (entity.ListMenuMasters, error) {
	menuMasters, err := g.menuMasterRepo.ListMenuMasters(ctx)
	if err != nil {
		return nil, err
	}
	return menuMasters, nil
}
