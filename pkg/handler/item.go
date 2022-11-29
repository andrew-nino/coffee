package handler

import (
	"coffee-app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {

}

type itemsById struct {
	Data []coffee.Items `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {

	category := c.Param("id")

	lists, err := h.services.CoffeeItem.GetItemsById(category)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemsById{
		Data: lists,
	})
}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
