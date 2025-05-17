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

func (h *orderHandler) GetUserOrders(c *gin.Context) {
	var req dto.GetOrdersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	orders, err := h.orderService.GetUserOrders(req.UserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, orders)
}

func (h *orderHandler) ChangeOrderStatus(c *gin.Context) {
	var req dto.ChangeOrderStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	if err := h.orderService.ChangeOrderStatus(req.OrderID, req.Status); errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}
