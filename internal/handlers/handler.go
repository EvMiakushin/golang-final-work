package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers interface {
	Register(router *gin.Engine)
}

type handler struct {
	Service
}

func NewHandler(service *Service) Handlers {
	return &handler{
		*service,
	}
}

func (h *handler) Register(r *gin.Engine) {
	r.GET("/", h.Handler)
}

func (h *handler) Handler(c *gin.Context) {
	result := h.Service.GetCacheData()

	response, err := json.MarshalIndent(result, " ", " ")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Data(http.StatusOK, "application/json; charset=utf-8", response)
}
