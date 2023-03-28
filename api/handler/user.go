package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUserRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().Create(&createUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.User().GetByID(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create user", http.StatusCreated, resp)
}

// Get By ID User godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id user", http.StatusBadRequest, "invalid user id")
		return
	}

	resp, err := h.storages.User().GetByID(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get user by id", http.StatusCreated, resp)
}

// Get List User godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list user", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list user", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.User().GetList(&models.GetListUserRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.user.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list user response", http.StatusOK, resp)
}

// Update User godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UpdateUser true "UpdateUserRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var updateUser models.UpdateUser

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id user", http.StatusBadRequest, "invalid user id")
		return
	}

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.handlerResponse(c, "update user", http.StatusBadRequest, err.Error())
		return
	}

	updateUser.Id = id

	rowsAffected, err := h.storages.User().Update(&updateUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.user.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.User().GetByID(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update user", http.StatusAccepted, resp)
}

// DELETE User godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UserPrimaryKey true "DeleteUserRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id user", http.StatusBadRequest, "invalid user id")
		return
	}

	rowsAffected, err := h.storages.User().Delete(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.user.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete user", http.StatusNoContent, nil)
}
