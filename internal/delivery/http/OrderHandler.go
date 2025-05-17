package http

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/services/order"
	"backnedTestGolang/pkg/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type OrderHandler interface {
	GetUserOrders(c *gin.Context)
	ChangeOrderStatus(c *gin.Context)
}

type orderHandler struct {
	orderService order.OrderService
}

func NewOrderHandler(orderService order.OrderService) OrderHandler {
	return &orderHandler{orderService: orderService}
}

// GetUserOrders godoc
// @Summary Get user's orders
// @Description Get a list of orders of a specific user
// @Tags orders
// @Accept json
// @Produce json
// @Param user_id query uint64 true "User ID"
// @Success 200 {object} dto.UserOrdersDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/get [get]
func (h *orderHandler) GetUserOrders(c *gin.Context) {
	var req dto.GetOrdersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	orders, err := h.orderService.GetUserOrders(req.UserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
	}

	c.JSON(http.StatusOK, orders)
}

// ChangeOrderStatus godoc
// @Summary Change order's status
// @Description changes the status of a certain order to the specified one
// @Tags orders
// @Accept json
// @Produce json
// @Param input body dto.ChangeOrderStatusRequest true "Order"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/change_status [post]
func (h *orderHandler) ChangeOrderStatus(c *gin.Context) {
	var req dto.ChangeOrderStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	if err := h.orderService.ChangeOrderStatus(req.OrderID, req.Status); errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
	}
	c.Status(http.StatusOK)
}
