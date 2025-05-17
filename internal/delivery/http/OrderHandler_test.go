package http

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOrderHandler_GetUserOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockOrderService)
	handler := NewOrderHandler(service)

	router := gin.New()

	router.GET("/order/get", handler.GetUserOrders)

	t.Run("valid request", func(t *testing.T) {
		expected := dto.UserOrdersDTO{UserID: 1, Orders: []dto.OrderDTO{{OrderID: 1, Status: "status", Items: make([]dto.ItemDTO, 0)}, {OrderID: 2, Status: "status", Items: make([]dto.ItemDTO, 0)}}}
		service.On("GetUserOrders", uint64(1)).Return(&expected, nil)

		req, _ := http.NewRequest(http.MethodGet, "/order/get?user_id=1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("invalid request", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/order/get", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		var empty *dto.UserOrdersDTO
		service.On("GetUserOrders", uint64(1)).Return(empty, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodGet, "/order/get?user_id=1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestOrderHandler_ChangeOrderStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockOrderService)
	handler := NewOrderHandler(service)

	router := gin.New()
	router.POST("/order/change_status", handler.ChangeOrderStatus)

	t.Run("valid request", func(t *testing.T) {
		service.On("ChangeOrderStatus", uint64(1), "status").Return(nil)

		body := `{"order_id": 1, "status": "status"}`
		req, _ := http.NewRequest(http.MethodPost, "/order/change_status", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("invalid request", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodPost, "/order/change_status", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		service.On("ChangeOrderStatus", uint64(1), "status").Return(errors.New("service error"))

		body := `{"order_id": 1, "status": "status"}`
		req, _ := http.NewRequest(http.MethodPost, "/order/change_status", strings.NewReader(body))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
