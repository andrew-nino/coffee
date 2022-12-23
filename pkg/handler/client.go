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
	Email                 string  `json:"email" db:"email"`
	Birthday              string  `json:"birthday" db:"birthday"`
	Sex                   string  `json:"sex" db:"sex"`
	IsAgreeToNotification bool    `json:"isAgreeToNotification" db:"isAgreeToNotification"`
	Comment               string  `json:"comment" db:"comment"`
	Value                 float32 `json:"value" db:"value"`
}

func (h *Handler) getBalance(c *gin.Context) {

	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	balanceFloat, err := h.services.CoffeeClient.GetBalance(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var balanceInt = int(balanceFloat)

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balanceInt,
	})
}
