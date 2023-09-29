package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
	menuMasterMock "github.com/kokopelli-inc/echo-ddd-demo/internal/domain/repository/mock"
	"github.com/kokopelli-inc/echo-ddd-demo/internal/infra/logger"
	dbError "github.com/kokopelli-inc/echo-ddd-demo/internal/infra/postgres"
)

const (
	testMenuMasterID = 1
)

func TestMenuMasterHandler_ListMenuMasters(t *testing.T) {
	a := assert.New(t)
	zapLogger, err := logger.NewLogger(true)
	a.NoError(err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMenuMasterRepo := menuMasterMock.NewMockMenuMasterRepository(ctrl)
	mockMenuMasterHandler := NewMenuMasterHandler(mockMenuMasterRepo, zapLogger)

	tests := []struct {
		id            int
		name          string
		branchID      string
		mockCall      bool
		expected      entity.ListMenuMasters
		expectedError dbError.DBError
		wantStatus    int
		wantBody      func() string
		wantError     error
		wantErrorBody map[string]any
	}{
		{
			id:            1,
			name:          "正常系",
			branchID:      "101",
			mockCall:      true,
			expected:      expectedListMenuMasters(),
			expectedError: nil,
			wantStatus:    http.StatusOK,
			wantBody: func() string {
				res := createListMenuMastersResponse()
				wantBody, _ := json.Marshal(res)
				return string(wantBody)
			},
		},
		{
			id:            2,
			name:          "異常系/DBエラー",
			branchID:      "101",
			mockCall:      true,
			expected:      nil,
			expectedError: errors.New("db error"),
			wantStatus:    http.StatusInternalServerError,
			wantError:     errors.New("code=500, message=db error"),
			wantBody: func() string {
				res := errResponse{
					Message: "db error",
					Status:  http.StatusInternalServerError,
				}
				wantBody, _ := json.Marshal(res)
				return string(wantBody)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			a = assert.New(t)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/branches/%s/menus", tt.branchID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("trace_id", "test_trace_id")
			c.SetPath("/branches/:branch_id/menus")
			c.SetParamNames("branch_id")
			c.SetParamValues(tt.branchID)

			// バリデーションエラーなどではmockを呼び出さないケースもある
			if tt.mockCall {
				mockMenuMasterRepo.EXPECT().ListMenuMasters(c.Request().Context()).Return(tt.expected, tt.expectedError)
			}
			err = mockMenuMasterHandler.ListMenuMasters(c)

			if tt.wantError != nil {
				a.Equal(tt.wantStatus, rec.Code)
				a.Equal(tt.wantBody(), strings.TrimSpace(rec.Body.String()))
				return
			}

			a.NoError(err)
			a.Equal(tt.wantStatus, rec.Code)
			// response bodyに改行コードが入るので取り除いてから比較
			a.Equal(tt.wantBody(), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func createListMenuMastersResponse() listMenuMasterResponse {
	return []menuMasterResponse{
		{
			ID:   1,
			Name: "相続相談",
		},
		{
			ID:   2,
			Name: "融資相談",
		},
	}
}

// ListMenuMastersのmockの期待値を作成
func expectedListMenuMasters() entity.ListMenuMasters {
	return entity.ListMenuMasters{
		{
			ID:   1,
			Name: "相続相談",
		},
		{
			ID:   2,
			Name: "融資相談",
		},
	}
}

func TestMenuMasterHandler_RegisterMenuMaster(t *testing.T) {
	a := assert.New(t)
	zapLogger, err := logger.NewLogger(true)
	a.NoError(err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMenuMasterRepo := menuMasterMock.NewMockMenuMasterRepository(ctrl)
	mockMenuMasterHandler := NewMenuMasterHandler(mockMenuMasterRepo, zapLogger)

	tests := []struct {
		id            int
		name          string
		request       *registerMenuMasterRequest
		mockCall      bool
		expected      int
		expectedError dbError.DBError
		wantStatus    int
		wantBody      func() string
		wantError     error
	}{
		{
			id:   1,
			name: "正常系",
			request: &registerMenuMasterRequest{
				BankID: "0158",
				Name:   "相続相談",
			},
			mockCall:      true,
			expected:      testMenuMasterID,
			expectedError: nil,
			wantStatus:    http.StatusCreated,
			wantBody: func() string {
				return createRegisterMenuMasterResponse()
			},
		},
		{
			id:         2,
			name:       "異常系/バリデーションエラー/主にrequired",
			request:    &registerMenuMasterRequest{},
			mockCall:   false,
			wantStatus: http.StatusBadRequest,
			wantBody: func() string {
				errRes := []fieldError{
					{
						Field: "銀行ID",
						Error: "required",
					},
					{
						Field: "メニューマスタ名",
						Error: "required",
					},
				}
				wantBody, _ := json.Marshal(errRes)
				return string(wantBody)
			},
		},
		{
			id:   3,
			name: "異常系/バリデーションエラー/主に文字数などの制約系",
			request: &registerMenuMasterRequest{
				BankID: "01589",
				Name:   strings.Repeat("a", 256),
			},
			mockCall:   false,
			wantStatus: http.StatusBadRequest,
			wantBody: func() string {
				errRes := []fieldError{
					{
						Field: "銀行ID",
						Error: "len",
					},
					{
						Field: "メニューマスタ名",
						Error: "max",
					},
				}
				wantBody, _ := json.Marshal(errRes)
				return string(wantBody)
			},
		},
		{
			id:   4,
			name: "異常系/DBエラー",
			request: &registerMenuMasterRequest{
				BankID: "0158",
				Name:   strings.Repeat("a", 15),
			},
			mockCall:      true,
			expectedError: &dbError.InternalServerError{Message: "DBエラー"},
			wantStatus:    http.StatusInternalServerError,
			wantBody: func() string {
				res := errResponse{
					Message: "DBエラー",
					Status:  http.StatusInternalServerError,
				}
				wantBody, _ := json.Marshal(res)
				return string(wantBody)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a = assert.New(t)
			e := echo.New()

			var req *http.Request
			jsonStr, _ := json.Marshal(tt.request)

			req = httptest.NewRequest(http.MethodPost, "/menu_masters", bytes.NewBuffer(jsonStr))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("trace_id", "test_trace_id")
			c.SetPath("/menu_masters")

			// バリデーションエラーなどではmockを呼び出さないケースもある
			if tt.mockCall {
				mockMenuMasterRepo.EXPECT().RegisterMenuMaster(c.Request().Context(), gomock.Any()).Return(tt.expected, tt.expectedError)
			}
			_ = mockMenuMasterHandler.RegisterMenuMaster(c)

			a.Equal(tt.wantStatus, rec.Code)
			// response bodyに改行コードが入るので取り除いてから比較
			a.Equal(tt.wantBody(), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func createRegisterMenuMasterResponse() string {
	resp := newCreatedResponse(testMenuMasterID)

	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}
