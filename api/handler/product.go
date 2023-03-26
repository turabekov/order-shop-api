package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {

	var createProduct models.CreateProduct

	err := c.ShouldBindJSON(&createProduct) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create product", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Product().Create(&createProduct)
	if err != nil {
		h.handlerResponse(c, "storage.product.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create product", http.StatusCreated, resp)
}

func (h *Handler) GetByIdProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id product", http.StatusBadRequest, "invalid product id")
		return
	}

	resp, err := h.storages.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get product by id", http.StatusCreated, resp)
}

func (h *Handler) GetListProduct(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Product().GetList(&models.GetListProductRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.product.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list product response", http.StatusOK, resp)
}

func (h *Handler) UpdateProduct(c *gin.Context) {

	var updateProduct models.UpdateProduct

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id product", http.StatusBadRequest, "invalid product id")
		return
	}

	err := c.ShouldBindJSON(&updateProduct)
	if err != nil {
		h.handlerResponse(c, "update product", http.StatusBadRequest, err.Error())
		return
	}

	updateProduct.Id = id

	rowsAffected, err := h.storages.Product().Update(&updateProduct)
	if err != nil {
		h.handlerResponse(c, "storage.product.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.product.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update product", http.StatusAccepted, resp)
}

func (h *Handler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id product", http.StatusBadRequest, "invalid product id")
		return
	}

	rowsAffected, err := h.storages.Product().Delete(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.product.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete product", http.StatusNoContent, nil)
}
