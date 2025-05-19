package http

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/mocks"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCartHandler_AddProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockCartService)
	handler := NewCartHandler(service)

	router := gin.New()
	router.POST("/cart/product/add", handler.AddProduct)

	t.Run("valid request", func(t *testing.T) {
		service.On("AddProduct", uint64(1), uint64(2), 3).Return(nil)
		body := `{"cart_id": 1, "product_id": 2, "quantity": 3}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/product/add", strings.NewReader(body))

		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		service.AssertCalled(t, "AddProduct", uint64(1), uint64(2), 3)
	})

	t.Run("validation error", func(t *testing.T) {
		body := `{"cart_id": 1, "product_id": 2}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/product/add", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		service.On("AddProduct", uint64(1), uint64(2), 3).Return(errors.New("service error"))

		body := `{"cart_id": 1, "product_id": 2, "quantity": 3}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/product/add", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestCartHandler_RemoveProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(mocks.MockCartService)
	handler := NewCartHandler(service)

	router := gin.New()
	router.POST("/cart/product/remove", handler.RemoveProduct)

	t.Run("valid request", func(t *testing.T) {
		service.On("RemoveProduct", uint64(1), uint64(2)).Return(nil)

		body := `{"cart_id": 1, "product_id": 2}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/product/remove", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("validation error", func(t *testing.T) {
		body := `{"cart_id": 1}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/product/remove", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		service.On("RemoveProduct", uint64(1), uint64(2)).Return(errors.New("service error"))

		body := `{"cart_id": 1, "product_id": 2}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/product/remove", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestCartHandler_ReduceProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(mocks.MockCartService)
	handler := NewCartHandler(service)

	router := gin.New()
	router.POST("/cart/reduce", handler.ReduceProduct)

	t.Run("valid request", func(t *testing.T) {
		service.On("ReduceProduct", uint64(1), uint64(2), 3).Return(nil)

		body := `{"cart_id": 1, "product_id": 2, "quantity": 3}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/reduce", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("validation error", func(t *testing.T) {
		body := `{"cart_id": 1}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/reduce", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		service.On("ReduceProduct", uint64(1), uint64(2), 3).Return(errors.New("service error"))

		body := `{"cart_id": 1, "product_id": 2, "quantity": 3}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/reduce", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestCartHandler_GetCartItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(mocks.MockCartService)
	handler := NewCartHandler(service)

	router := gin.New()
	router.GET("/cart/get", handler.GetCartItems)

	t.Run("valid request", func(t *testing.T) {
		expected := []dto.ItemDTO{
			{Name: "Item A", Price: 100, Quantity: 2},
			{Name: "Item B", Price: 50, Quantity: 1},
		}

		service.On("GetCartItems", uint64(1)).Return(&expected, nil)

		req, _ := http.NewRequest(http.MethodGet, "/cart/get?user_id=1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var actual []dto.ItemDTO
		err := json.Unmarshal(resp.Body.Bytes(), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("validation error", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/cart/get", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		var empty *[]dto.ItemDTO
		service.On("GetCartItems", uint64(1)).Return(empty, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodGet, "/cart/get?user_id=1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestCartHandler_Checkout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockCartService)
	handler := NewCartHandler(service)

	router := gin.New()
	router.POST("/cart/checkout", handler.Checkout)

	t.Run("valid request", func(t *testing.T) {
		service.On("Checkout", uint64(1)).Return(nil)

		body := `{"user_id": 1}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/checkout", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		service.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/cart/checkout", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("service error", func(t *testing.T) {
		service.On("Checkout", uint64(1)).Return(errors.New("service error"))

		body := `{"user_id": 1}`
		req, _ := http.NewRequest(http.MethodPost, "/cart/checkout", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
