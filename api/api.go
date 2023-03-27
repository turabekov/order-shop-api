package api

import (
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)

	// customer api
	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	// user api
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	// courier api
	r.POST("/courier", handler.CreateCourier)
	r.GET("/courier/:id", handler.GetByIdCourier)
	r.GET("/courier", handler.GetListCourier)
	r.PUT("/courier/:id", handler.UpdateCourier)
	r.DELETE("/courier/:id", handler.DeleteCourier)

	// category api
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// product api
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	// order api
	r.POST("/order", handler.CreateOrder)
	r.GET("/order/:id", handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrder)
	r.PUT("/order/:id", handler.UpdateOrder)
	r.PATCH("/order/:id", handler.UpdateOrder)
	r.DELETE("/order/:id", handler.DeleteOrder)
}
