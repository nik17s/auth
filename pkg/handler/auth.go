package handler

import (
	"github.com/gin-gonic/gin"
	entity "github.com/nik17s/auth"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var user entity.User

	err := c.BindJSON(&user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.CreateUser(user.Login, user.Password, user.Email, user.Phone)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var user entity.User
	err := c.BindJSON(&user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.GenerateJWT(user.Login, user.Password)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"jwt": token,
	})

}
