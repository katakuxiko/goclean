package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katakuxiko/clean_go/structure"
)

func (h *Handler) createUserVariables(c *gin.Context) {
	userId, err := getUserId(c)
	if (err != nil) {
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	var input structure.UsersVariables;
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	id,err := h.service.User.Create(userId,input)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"id":id,
	})
}