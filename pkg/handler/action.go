package handler

import (
	"coffee-app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Actions struct {
	Data []coffee.Action `json:"data"`
}

func (h *Handler) getAllActions(c *gin.Context) {

	lists, err := h.services.CoffeeAction.GetActions()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Actions{
		Data: lists,
	})
}

func (h *Handler) getActionByID(c *gin.Context) {

	guid := c.Param("id")

	action, err := h.services.CoffeeAction.GetActionById(guid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, action)
}
