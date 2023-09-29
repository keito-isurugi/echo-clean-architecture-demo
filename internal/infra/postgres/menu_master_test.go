package postgres_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/db"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/env"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/logger"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/postgres"
)

func TestMenuMasterRepository_ListMenuMaster(t *testing.T) {
	tests := []struct {
		id    int
		name  string
		want  entity.ListMenuMasters
		setup func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/一覧取得",
			want: entity.ListMenuMasters{
				{
					ID:        1,
					BankID:    "0158",
					Name:      "相続相談",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					CreatedBy: "admin",
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedBy: "admin",
				},
				{
					ID:        2,
					BankID:    "0158",
					Name:      "融資相談",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					CreatedBy: "admin",
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedBy: "admin",
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO banks (id, name, created_by, updated_by) VALUES ('0158', '京都銀行', 'admin', 'admin')`).Error)

				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO bap_test.public.menu_masters (id, bank_id, name, created_at, created_by, updated_at, updated_by) VALUES (1, '0158', '相続相談', '2021-01-01T00:00:00+00:00', 'admin', '2021-01-01T00:00:00+00:00', 'admin')`).Error)

				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO bap_test.public.menu_masters (id, bank_id, name, created_at, created_by, updated_at, updated_by) VALUES (2, '0158', '融資相談', '2021-01-01T00:00:00+00:00', 'admin', '2021-01-01T00:00:00+00:00', 'admin')`).Error)
			},
		},
		{
			id:   2,
			name: "対象のレコードがない場合、空の配列が返ってくること",
			want: entity.ListMenuMasters{},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			zapLogger, err := logger.NewLogger(true)
			a.NoError(err)

			ev, err := env.NewValue()
			a.NoError(err)

			dbClient, err := db.NewClient(&ev.DB, zapLogger)
			a.NoError(err)

			truncateTable(ctx, t, dbClient)

			if tt.setup != nil {
				tt.setup(ctx, t, dbClient)
			}

			menuMasterRepo := postgres.NewMenuMasterRepository(dbClient, zapLogger)

			got, err := menuMasterRepo.ListMenuMasters(ctx)
			a.NoError(err)

			if tt.want != nil {
				if !cmp.Equal(got, tt.want) {
					t.Errorf("diff %s", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}

func TestMenuMasterRepository_RegisterMenuMaster(t *testing.T) {
	tests := []struct {
		id        int
		name      string
		request   *entity.MenuMaster
		wantTable *entity.MenuMaster
		wantError error
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/新規作成",
			request: &entity.MenuMaster{
				BankID:    "0158",
				Name:      "投資相談",
				CreatedBy: "admin",
				UpdatedBy: "admin",
			},
			wantTable: &entity.MenuMaster{
				ID:        1,
				BankID:    "0158",
				Name:      "投資相談",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				CreatedBy: "admin",
				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedBy: "admin",
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO banks (id, name, created_by, updated_by) VALUES ('0158', '京都銀行', 'admin', 'admin')`).Error)
			},
		},
		{
			id:   2,
			name: "異常系/レコード重複",
			request: &entity.MenuMaster{
				BankID:    "0158",
				Name:      "投資相談",
				CreatedBy: "admin",
				UpdatedBy: "admin",
			},
			wantError: &postgres.DuplicateError{
				Message: "ERROR: duplicate key value violates unique constraint \"idx_menu_masters_bank_id_name\" (SQLSTATE 23505)",
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO banks (id, name, created_by, updated_by) VALUES ('0158', '京都銀行', 'admin', 'admin')`).Error)

				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO menu_masters (bank_id, name, created_at,created_by,updated_at, updated_by) VALUES ('0158', '投資相談', '2021-01-01T00:00:00+00:00','admin','2021-01-01T00:00:00+00:00', 'admin');`).Error)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			zapLogger, err := logger.NewLogger(true)
			a.NoError(err)

			ev, err := env.NewValue()
			a.NoError(err)

			dbClient, err := db.NewClient(&ev.DB, zapLogger)
			a.NoError(err)

			truncateTable(ctx, t, dbClient)

			if tt.setup != nil {
				tt.setup(ctx, t, dbClient)
			}

			menuMasterRepo := postgres.NewMenuMasterRepository(dbClient, zapLogger)
			_, err = menuMasterRepo.RegisterMenuMaster(ctx, tt.request)

			if tt.wantError != nil {
				a.EqualError(err, tt.wantError.Error())
				return
			}

			a.NoError(err)

			var got *entity.MenuMaster
			if err = dbClient.Conn(ctx).Find(&got).Error; err != nil {
				a.EqualError(err, tt.wantError.Error())
				return
			}

			if tt.wantTable != nil {
				// ID, CreatedAt, UpdatedAtはランダムの値が入ってくるので比較対象外
				opt := cmpopts.IgnoreFields(entity.MenuMaster{}, "ID", "CreatedAt", "UpdatedAt")
				if !cmp.Equal(got, tt.wantTable, opt) {
					t.Errorf("diff %s", cmp.Diff(got, tt.wantTable))
				}
			}
		})
	}
}
