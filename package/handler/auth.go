package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katakuxiko/clean_go/structure"
)

func (h *Handler) signUp(c *gin.Context) {
	var input structure.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	id,err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id" : id,
	})
}

type singInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	token,err := h.service.Authorization.GenerateToken(input.Username,input.Password)
	if err != nil {
		NewErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token" : token,
	})
}