package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type clienFromYtimes struct {
	EventId                    string  `json:"eventId,omitempty"`
	DateTimeUTC                string  `json:"dateTimeUTC,omitempty"`
	PphoneCode                 string  `json:"phoneCode,omitempty"`
	Phone                      string  `json:"phone,omitempty"`
	Number                     int     `json:"number,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Surname                    string  `json:"surname,omitempty"`
	Email                      string  `json:"email,omitempty"`
	Birthday                   string  `json:"birthday,omitempty"`
	Sex                        string  `json:"sex,omitempty"`
	Comment                    string  `json:"comment,omitempty"`
	IsAgreeToNotification      bool    `json:"isAgreeToNotification,omitempty"`
	ExcludeAddPoints           bool    `json:"excludeAddPoints,omitempty"`
	PointsValue                float64 `json:"pointsValue,omitempty"`
	PointsChange               float64 `json:"pointsChange,omitempty"`
	PointsChangeComment        string  `json:"pointsChangeComment,omitempty"`
	StatVisitCount             int     `json:"statVisitCount,omitempty"`
	StatPayValue               float64 `json:"statPayValue,omitempty"`
	StatLastMonthVisitCount    int     `json:"statLastMonthVisitCount,omitempty"`
	StatLastMonthPayValue      float64 `json:"statLastMonthPayValue,omitempty"`
	StatCurrentMonthVisitCount int     `json:"statCurrentMonthVisitCount,omitempty"`
	StatCurrentMonthPayValue   float64 `json:"statCurrentMonthPayValue,omitempty"`
}

func (h *Handler) updateClient(c *gin.Context) {

	var input clienFromYtimes

	if err := c.BindJSON(&input); err != nil {

		logrus.Error("parsing error")
		return
	}

	fmt.Println(input)

	c.String(http.StatusOK, "OK")
}
