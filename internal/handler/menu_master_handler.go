package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	domain "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository"
	usecase "github.com/kokopelli-inc/echo-ddd-demo/internal/usecase/menu_master"
)

type MenuMasterHandler interface {
	ListMenuMasters(echo.Context) error
	RegisterMenuMaster(echo.Context) error
}

type menuMasterResponse struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"相続相談"`
}

type listMenuMasterResponse []menuMasterResponse

type menuMasterHandler struct {
	menuMasterRepo domain.MenuMasterRepository
	zapLogger      *zap.Logger
}

func NewMenuMasterHandler(menuMasterRepo domain.MenuMasterRepository, zapLogger *zap.Logger) MenuMasterHandler {
	return &menuMasterHandler{
		menuMasterRepo: menuMasterRepo,
		zapLogger:      zapLogger,
	}
}

// ListMenuMasters
// @Summary		メニューマスタ一覧
// @Description	メニューマスタ一覧を取得
// @id			ListMenuMasters
// @tags		menu_master
// @Accept		json
// @Produce		json
// @Success		200	{object}	listMenuMasterResponse
// @Failure		401	{object}	errResponse
// @Failure		403	{object}	errResponse
// @Failure		500	{object}	errResponse
// @Router		/menu_masters [get]
// @Param		Authorization	header	string	true	"Bearer {token}"
// @Param		REALM							header	string	true	"bank0158"
// @Param		X-BA-PORTAL-USER-UNIQUE-CODE	header	string	true	"test"
func (a *menuMasterHandler) ListMenuMasters(c echo.Context) error {
	traceID := c.Get("trace_id").(string)
	uc := usecase.NewListMenuMasterUsecase(a.menuMasterRepo)
	menuMasters, err := uc.Exec(c.Request().Context())
	if err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(a.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}

	res := make(listMenuMasterResponse, len(menuMasters))
	for i, menuMaster := range menuMasters {
		res[i] = menuMasterResponse{
			ID:   menuMaster.ID,
			Name: menuMaster.Name,
		}
	}
	return c.JSON(http.StatusOK, res)
}

type registerMenuMasterRequest struct {
	BankID string `json:"bank_id" example:"0158" ja:"銀行ID" validate:"required,len=4"`
	Name   string `json:"name" example:"相続相談" ja:"メニューマスタ名" validate:"required,max=255"`
}

// RegisterMenuMaster
// @Summary		メニューマスタ登録
// @Description	メニューマスタを登録
// @id			RegisterMenuMaster
// @tags		menu_master
// @Accept		json
// @Produce		json
// @Success		201			{object}	createdResponse
// @Failure		400			{object}	fieldError
// @Failure		401			{object}	errResponse
// @Failure		403			{object}	errResponse
// @Failure		500			{object}	errResponse
// @Router		/menu_masters [post]
// @Param		Authorization	header	string	true	"Bearer {token}"
// @Param		REALM							header	string	true	"bank0158"
// @Param		X-BA-PORTAL-USER-UNIQUE-CODE	header	string	true	"test"
// @Param request body registerMenuMasterRequest true "registerMenuMasterRequest"
func (a *menuMasterHandler) RegisterMenuMaster(c echo.Context) error {
	traceID := c.Get("trace_id").(string)

	var req registerMenuMasterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(req); err != nil {
		a.zapLogger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, fieldErrors(err))
	}

	uc := usecase.NewRegisterMenuMasterUsecase(a.menuMasterRepo)
	newMenuMaster := entity.NewRegisterMenuMaster(
		req.BankID,
		req.Name,
		"admin",
		"admin",
	)

	id, err := uc.Exec(c.Request().Context(), newMenuMaster)
	if err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(a.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}

	return c.JSON(http.StatusCreated, createdResponse{ID: id})
}