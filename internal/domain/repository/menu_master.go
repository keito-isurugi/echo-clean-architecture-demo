//go:generate mockgen -source=menu_master.go -destination=./mock/menu_master_mock.go
package domain

import (
	"context"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
)

type MenuMasterRepository interface {
	ListMenuMasters(ctx context.Context) (entity.ListMenuMasters, error)
	RegisterMenuMaster(ctx context.Context, menuMaster *entity.MenuMaster) (int, error)
}
