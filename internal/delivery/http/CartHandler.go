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

// AddProduct godoc
// @Summary Add a product to the cart
// @Description Adds a product with quantity to the cart
// @Tags cart
// @Accept json
// @Produce json
// @Param input body dto.AddProductRequest true "Cart Product"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/product/add [post]
func (h *cartHandler) AddProduct(c *gin.Context) {
	var req dto.AddProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	err := h.cartService.AddProduct(req.CartID, req.ProductID, req.Quantity)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "product not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

// ReduceProduct godoc
// @Summary Reduces a product in the cart
// @Description Reduces the product by the specified quantity in the cart
// @Tags cart
// @Accept json
// @Produce json
// @Param input body dto.ReduceProductRequest true "Cart Product"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/product/reduce [post]
func (h *cartHandler) ReduceProduct(c *gin.Context) {
	var req dto.ReduceProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	err := h.cartService.ReduceProduct(req.CartID, req.ProductID, req.Quantity)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "product not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}
	c.Status(http.StatusOK)
}

// RemoveProduct godoc
// @Summary Removes product from the cart
// @Description Reduces the product by the specified quantity in the cart
// @Tags cart
// @Accept json
// @Produce json
// @Param input body dto.ReduceProductRequest true "Cart Product"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/product/remove [post]
func (h *cartHandler) RemoveProduct(c *gin.Context) {
	var req dto.RemoveProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	err := h.cartService.RemoveProduct(req.CartID, req.ProductID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}
	c.Status(http.StatusOK)
}

// GetCartItems godoc
// @Summary Get items in user's cart
// @Description Get a list of items from a specific cart
// @Tags cart
// @Accept json
// @Produce json
// @Param user_id query uint64 true "User ID"
// @Success 200 {array} []dto.ItemDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/get [get]
func (h *cartHandler) GetCartItems(c *gin.Context) {
	var req dto.GetCartRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	cartItems, err := h.cartService.GetCartItems(req.UserID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(fmt.Errorf("record not found"))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "record not found"})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}

// Checkout godoc
// @Summary Check out an order
// @Description Check out an order with the contents of the cart, clearing the cart
// @Tags cart
// @Accept json
// @Produce json
// @Param input body dto.CheckoutRequest true "User ID"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/checkout [post]
func (h *cartHandler) Checkout(c *gin.Context) {
	var req dto.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid input", Details: logger.ValidationErrors(err).Error()})
		return
	}

	if err := h.cartService.Checkout(req.UserID); err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}
	c.Status(http.StatusOK)
}
