package http

import (
	"backnedTestGolang/internal/dto"
	"backnedTestGolang/internal/services/cart"
	"backnedTestGolang/pkg/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CartHandler interface {
	AddProduct(c *gin.Context)
	ReduceProduct(c *gin.Context)
	RemoveProduct(c *gin.Context)
	GetCartItems(c *gin.Context)
	Checkout(c *gin.Context)
}

type cartHandler struct {
	cartService cart.CartService
}

func NewCartHandler(cartService cart.CartService) CartHandler {
	return &cartHandler{cartService: cartService}
}

func (h *cartHandler) AddProduct(c *gin.Context) {
	var req dto.AddProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	err := h.cartService.AddProduct(req.CartID, req.ProductID, req.Quantity)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *cartHandler) ReduceProduct(c *gin.Context) {
	var req dto.ReduceProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	err := h.cartService.ReduceProduct(req.CartID, req.ProductID, req.Quantity)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *cartHandler) RemoveProduct(c *gin.Context) {
	var req dto.RemoveProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	err := h.cartService.RemoveProduct(req.CartID, req.ProductID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *cartHandler) GetCartItems(c *gin.Context) {
	var req dto.GetCartRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	cartItems, err := h.cartService.GetCartItems(req.UserID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}

func (h *cartHandler) Checkout(c *gin.Context) {
	var req dto.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": logger.ValidationErrors(err).Error()})
		return
	}

	if err := h.cartService.Checkout(req.UserID); err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
