package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/Register", h.Register)
		auth.POST("/Login", h.Login)
	}

	return router
}
