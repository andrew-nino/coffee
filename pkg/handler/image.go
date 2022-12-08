package handler

import (
	"github.com/gin-gonic/gin"
)

const DIRECTORY = "image"

func (h *Handler) getImage(c *gin.Context) {

	nameImage := c.Param("name")

	c.Set("Content-Type", "image/jpeg")

	c.File(DIRECTORY + "/" + nameImage)
}
