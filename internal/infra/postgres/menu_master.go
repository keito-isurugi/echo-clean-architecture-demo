package postgres

import (
	"context"

	"go.uber.org/zap"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
)

type menuMasterRepository struct {
	dbClient  db.Client
	zapLogger *zap.Logger
}

func NewMenuMasterRepository(dbClient db.Client, zapLogger *zap.Logger) domain.MenuMasterRepository {
	return &menuMasterRepository{
		dbClient:  dbClient,
		zapLogger: zapLogger,
	}
}

func (r *menuMasterRepository) ListMenuMasters(ctx context.Context) (entity.ListMenuMasters, error) {
	var menuMasters entity.ListMenuMasters
	if err := r.dbClient.Conn(ctx).
		Find(&menuMasters).Error; err != nil {
		return nil, err
	}
	return menuMasters, nil
}

func (a *menuMasterRepository) RegisterMenuMaster(ctx context.Context, menuMaster *entity.MenuMaster) (int, error) {
	if err := a.dbClient.Conn(ctx).Create(&menuMaster).Error; err != nil {
		return 0, err
	}
	return menuMaster.ID, nil
}
