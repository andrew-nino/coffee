package handler

import (
	"coffee-app/assets"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getImage(c *gin.Context) {

	nameImage := c.Param("name")

	c.Set("Content-Type", "image/jpeg")

	c.File(assets.IMAGES + "/" + nameImage)
}
