package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type clientFromYTimes struct {
	EventId                    string  `json:"eventId,omitempty"`
	DateTimeUTC                string  `json:"dateTimeUTC,omitempty"`
	PhoneCode                  string  `json:"phoneCode,omitempty"`
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

type DataToSend struct {
	MessageKey   string       `json:"to"`
	Notification Notification `json:"notification"`
	Data         Data         `json:"data"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Data struct {
	ClickAction string  `json:"click_action"`
	Amount      float32 `json:"amount"`
}

func (h *Handler) updateClient(c *gin.Context) {

	var input clientFromYTimes

	if err := c.BindJSON(&input); err != nil {

		logrus.Error("parsing error")
		return
	}
	fmt.Println(input)
	c.String(http.StatusOK, "OK")
	
	fmt.Println("Request on DB")
	points, err := h.services.CoffeeDBUpdate.UpdatePoints(input.Phone, float32(input.PointsChange))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Answer from DB  Points = ", points)
	fmt.Println("Request Push")
	err = pushRequest(points)
	if err != nil {
		logrus.Error("pushRequest error")
	}
}

func (h *Handler) updateMenu(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}

func pushRequest(points float32) error {

	newData := DataToSend{
		MessageKey: "e1pTCS4vREWTP0syI8wQyf:APA91bF6KNwD6UzO7bKuTPU87nzEQY1aJjEaj41_bMUp7uovqHTJTSI1GXSnX23hTeGRToCkiRloQxn40IvGZf4slJNc_23Fgvn7-ptQcuFG2PDQDz33kPJIA_BMhDLz-XXSrAe2JOX7",
		Notification: Notification{
			Title: "hello",
			Body:  "hello body",
		},
		Data: Data{
			ClickAction: "FLUTTER_NOTIFICATION_CLICK",
			Amount:      points,
		},
	}

	reqjson, err := json.Marshal(newData)

	req, err := http.NewRequest(
		"POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(reqjson))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	respons, err := client.Do(req)
	if err != nil {
		return err
	}
	defer respons.Body.Close()

	fmt.Println("PUSH Respons Status = ", respons.Status)
	fmt.Println("PUSH Respons Body = ", respons.Body)

	return nil
}
