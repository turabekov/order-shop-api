package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrder(c *gin.Context) {

	var createOrder models.CreateOrder

	err := c.ShouldBindJSON(&createOrder) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create order", http.StatusBadRequest, err.Error())
		return
	}

	// check existing of product
	product, err := h.storages.Product().GetByID(&models.ProductPrimaryKey{Id: createOrder.ProductId})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// check existing of user
	_, err = h.storages.User().GetByID(&models.UserPrimaryKey{Id: createOrder.UserId})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// check existing of customer
	_, err = h.storages.Customer().GetByID(&models.CustomerPrimaryKey{Id: createOrder.CustomerId})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// calc total price
	createOrder.Price = float64(createOrder.Quantity) * product.Price

	id, err := h.storages.Order().Create(&createOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Order().GetByID(&models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create order", http.StatusCreated, resp)
}

func (h *Handler) GetByIdOrder(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id order", http.StatusBadRequest, "invalid order id")
		return
	}

	resp, err := h.storages.Order().GetByID(&models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get order by id", http.StatusCreated, resp)
}

func (h *Handler) GetListOrder(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Order().GetList(&models.GetListOrderRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.order.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list order response", http.StatusOK, resp)
}

func (h *Handler) UpdateOrder(c *gin.Context) {

	var updateOrder models.UpdateOrder

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id order", http.StatusBadRequest, "invalid order id")
		return
	}

	err := c.ShouldBindJSON(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "update order", http.StatusBadRequest, err.Error())
		return
	}

	updateOrder.Id = id

	// check existing of product
	product, err := h.storages.Product().GetByID(&models.ProductPrimaryKey{Id: updateOrder.ProductId})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// check existing of user
	_, err = h.storages.User().GetByID(&models.UserPrimaryKey{Id: updateOrder.UserId})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// check existing of customer
	_, err = h.storages.Customer().GetByID(&models.CustomerPrimaryKey{Id: updateOrder.CustomerId})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// calc total price
	updateOrder.Price = float64(updateOrder.Quantity) * product.Price

	rowsAffected, err := h.storages.Order().Update(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByID(&models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update  order", http.StatusAccepted, resp)
}

// ------------PATCH----------------
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id order", http.StatusBadRequest, "invalid order id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch order", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Order().UpdatePatch(&object)
	if err != nil {
		h.handlerResponse(c, "storage.order.updatepatch", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.updatepatch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByID(&models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch order", http.StatusAccepted, resp)
}

func (h *Handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id order", http.StatusBadRequest, "invalid order id")
		return
	}

	rowsAffected, err := h.storages.Order().Delete(&models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete order", http.StatusNoContent, nil)
}
