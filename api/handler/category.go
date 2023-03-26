package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create category", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Category().Create(&createCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create category", http.StatusCreated, resp)
}

func (h *Handler) GetByIdCategory(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id category", http.StatusBadRequest, "invalid category id")
		return
	}

	resp, err := h.storages.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get category by id", http.StatusCreated, resp)
}

func (h *Handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list category", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list category", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Category().GetList(&models.GetListCategoryRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.category.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list category response", http.StatusOK, resp)
}

func (h *Handler) UpdateCategory(c *gin.Context) {

	var updateCategory models.UpdateCategory

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id category", http.StatusBadRequest, "invalid category id")
		return
	}

	err := c.ShouldBindJSON(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "update category", http.StatusBadRequest, err.Error())
		return
	}

	updateCategory.Id = id

	rowsAffected, err := h.storages.Category().Update(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.category.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update category", http.StatusAccepted, resp)
}

func (h *Handler) DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id category", http.StatusBadRequest, "invalid user id")
		return
	}

	rowsAffected, err := h.storages.Category().Delete(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.category.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete category", http.StatusNoContent, nil)
}
