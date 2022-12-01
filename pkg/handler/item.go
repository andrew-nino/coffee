package handler

import (
	"coffee-app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type itemsById struct {
	Data []coffee.Item `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {

	lists, err := h.services.CoffeeItem.GetItems()

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemsById{
		Data: lists,
	})
}

func (h *Handler) getItemFromCategory(c *gin.Context) {

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
