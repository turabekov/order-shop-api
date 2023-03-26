package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourier(c *gin.Context) {

	var createCourier models.CreateCourier

	err := c.ShouldBindJSON(&createCourier) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create courier", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Courier().Create(&createCourier)
	if err != nil {
		h.handlerResponse(c, "storage.courier.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Courier().GetByID(&models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create courier", http.StatusCreated, resp)
}

func (h *Handler) GetByIdCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	resp, err := h.storages.Courier().GetByID(&models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get courier by id", http.StatusCreated, resp)
}

func (h *Handler) GetListCourier(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list courier", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list courier", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Courier().GetList(&models.GetListCourierRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list courier response", http.StatusOK, resp)
}

func (h *Handler) UpdateCourier(c *gin.Context) {

	var updateCourier models.UpdateCourier

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	err := c.ShouldBindJSON(&updateCourier)
	if err != nil {
		h.handlerResponse(c, "update courier", http.StatusBadRequest, err.Error())
		return
	}

	updateCourier.Id = id

	rowsAffected, err := h.storages.Courier().Update(&updateCourier)
	if err != nil {
		h.handlerResponse(c, "storage.user.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.courier.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Courier().GetByID(&models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update courier", http.StatusAccepted, resp)
}

func (h *Handler) DeleteCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	rowsAffected, err := h.storages.Courier().Delete(&models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.courier.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete courier", http.StatusNoContent, nil)
}
