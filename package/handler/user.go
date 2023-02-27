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

type getAllVariablesResponse struct {
	Data structure.UsersVariables `json:"data"`
}

func (h *Handler) updateVariablesUser(c *gin.Context) {
	userId,err := getUserId(c)
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"invalid id param")
		return
	}
	
	var input structure.UpdateUserVariables
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.User.Update(userId,input)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) getAllVariables(c *gin.Context){
	userId,err := getUserId(c)
	
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	variables, err := h.service.User.GetAllVariables(userId)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllVariablesResponse{
		Data: variables,
	})
}