package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/katakuxiko/clean_go/structure"
)

func (h *Handler) createList(c *gin.Context) {
	userId,err := getUserId(c)
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	
	var input structure.BooksList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	
	id,err := h.service.BooksList.Create(userId,input)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"id":id,
	})
}

type getAllListsResponse struct {
	Data []structure.BooksList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId,err := getUserId(c)
	pageParam := c.Query("limit")
	if pageParam == "" {
		pageParam ="5"
	}

	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	lists,err := h.service.BooksList.GetAll(userId, pageParam)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getUserBooksAll(c *gin.Context) {
	userId,err := getUserId(c)
	
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	lists,err := h.service.BooksList.GetUserBooksAll(userId)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllListsResponse{
		Data: lists,
	})
}


func (h *Handler) getListById(c *gin.Context) {
	userId,err := getUserId(c)
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"invalid id param")
		return
	}
	list,err := h.service.BooksList.GetById(userId,id)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId,err := getUserId(c)
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"invalid id param")
		return
	}
	
	var input structure.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.BooksList.Update(userId,id,input)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId,err := getUserId(c)
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c,http.StatusInternalServerError,"invalid id param")
		return
	}
	err = h.service.BooksList.Delete(userId,id)
	if err != nil {
		NewErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,StatusResponse{
		Status: "ok",
	})
}