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
	Data []coffee.Category `json:"data"`
}

func (h *Handler) getLists(c *gin.Context) {

	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	category, _ := c.GetQuery("class")

	lists, err := h.services.CoffeeList.GetCategories(category)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getALLCategories{
		Data: lists,
	})
}
