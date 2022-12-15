package handler

import (
	"coffee-app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getTypes struct {
	Data []coffee.Type `json:"data"`
}

func (h *Handler) getTypesFromItem(c *gin.Context) {

	item := c.Param("id")

	lists, err := h.services.CoffeeTypes.GetTypes(item)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTypes{
		Data: lists,
	})
}
