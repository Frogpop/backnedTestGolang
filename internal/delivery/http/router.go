package http

import (
	"backnedTestGolang/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

type Router struct {
	cartHandler  CartHandler
	orderHandler OrderHandler
}

func NewRouter(cartHandler CartHandler, orderHandler OrderHandler) *Router {
	return &Router{cartHandler: cartHandler, orderHandler: orderHandler}
}

func (r *Router) Init(log *slog.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		middleware.RequsetLogger(log),
	)

	router.POST("/cart/product/add", r.cartHandler.AddProduct)
	router.POST("/cart/product/reduce", r.cartHandler.ReduceProduct)
	router.POST("/cart/product/remove", r.cartHandler.RemoveProduct)
	router.POST("/cart/checkout", r.cartHandler.Checkout)
	router.GET("/cart/get", r.cartHandler.GetCartItems)
	router.GET("/order/get", r.orderHandler.GetUserOrders)
	router.POST("/order/change_status", r.orderHandler.ChangeOrderStatus)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
