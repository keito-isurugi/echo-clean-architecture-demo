package usecase

import (
	"context"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
)

type RegisterMenuMastersUsecase interface {
	Exec(ctx context.Context, menuMaster *entity.MenuMaster) (int, error)
}

type registerMenuMastersUsecaseImpl struct {
	menuMasters domain.MenuMasterRepository
}

func NewRegisterMenuMasterUsecase(menuMasters domain.MenuMasterRepository) RegisterMenuMastersUsecase {
	return &registerMenuMastersUsecaseImpl{
		menuMasters: menuMasters,
	}
}

func (g *registerMenuMastersUsecaseImpl) Exec(
	ctx context.Context,
	menuMaster *entity.MenuMaster,
) (int, error) {
	id, err := g.menuMasters.RegisterMenuMaster(ctx, menuMaster)
	if err != nil {
		return 0, err
	}
	return id, nil
}
