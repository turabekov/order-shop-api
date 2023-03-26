package handler

import (
	"app/config"
	"app/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg      *config.Config
	storages storage.StorageI
}

type Response struct {
	Status      int
	Description string
	Data        interface{}
}

func NewHandler(cfg *config.Config, store storage.StorageI) *Handler {
	return &Handler{
		cfg:      cfg,
		storages: store,
	}
}

func (h *Handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	response := Response{
		Status:      code,
		Description: path,
		Data:        message,
	}

	c.JSON(code, response)
}

func (h *Handler) getOffsetQuery(offset string) (int, error) {
	if len(offset) <= 0 {
		return h.cfg.DefaultOffset, nil
	}

	return strconv.Atoi(offset)
}

func (h *Handler) getLimitQuery(limit string) (int, error) {

	if len(limit) <= 0 {
		return h.cfg.DefaultLimit, nil
	}

	return strconv.Atoi(limit)
}
