package handler

import (
	"coffee-app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {

	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getALLCategories struct {
	Data []coffee.Categories `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {

	lists, err := h.services.CoffeeList.GetALLCategories()

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getALLCategories{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
