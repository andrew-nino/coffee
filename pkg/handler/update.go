package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var g_counter int

func (h *Handler) updateDB(c *gin.Context) {

	if g_counter >= 1 {
		c.JSON(http.StatusMethodNotAllowed, map[string]interface{}{
			"rejected": "The update has already been completed",
		})
		return
	}

	data, err := h.services.CoffeeDBUpdate.UpdateDB()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)

	g_counter++
}
