package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CliientUpdate struct {
	RequestId             string  `json:"requestId"`
	PhoneCode             string  `json:"phoneCode" binding:"required" db:"phone_code"`
	Phone                 string  `json:"phone" binding:"required" db:"phone"`
	Name                  string  `json:"name" db:"name"`
	Surname               string  `json:"surname" db:"surname"`
	Email                 string  `json:"email" db:"email""`
	Birthday              string  `json:"birthday" db:"birthday"`
	Sex                   string  `json:"sex" db:"sex"`
	IsAgreeToNotification bool    `json:"isAgreeToNotification" db:"isAgreeToNotification"`
	Comment               string  `json:"comment" db:"comment"`
	Value                 float32 `json:"value" db:"value"`
}

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
