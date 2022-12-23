package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

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
	Amount      int `json:"amount"`
}

func (h *Handler) whClient(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty ayth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[1] != sendersUUID {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	var input clientFromYTimes

	if err := c.BindJSON(&input); err != nil {

		logrus.Error("parsing error")
		return
	}
	c.String(http.StatusOK, "OK")

	userData, err := h.services.CoffeeDBUpdate.UpdatePoints(input.Phone, float32(input.PointsValue))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var outValue =  int(userData.Value)

	err = pushRequest(outValue, userData.MessageKey)
	if err != nil {
		logrus.Error("pushRequest error")
	}
}

func (h *Handler) whMenu(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}

func pushRequest(points int, messageKey string) error {

	newData := DataToSend{
		MessageKey: messageKey,
		Notification: Notification{
			Title: "Спасибо за покупку!",
			Body:  "Текущий баланс: " + fmt.Sprint(points),
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
	req.Header.Add(authorizationHeader, "key=AAAAEgTcaiQ:APA91bEyrZpOVYoUD-lQRqxzS_zrzLn5WD-WQ3AtH-uvNNZQnF8ghT-_BaS0is5ptYS89vfAs9_34o2lr0I9abJ6dx3A7S2w1kKNQWJPzpR9c3o-4jg0ty0sxi3-0LlsDsYUqA_7yEQR")

	client := &http.Client{}
	respons, err := client.Do(req)
	if err != nil {
		return err
	}
	defer respons.Body.Close()

	if respons.Status != "200 OK" {
		responseBody, _ := ioutil.ReadAll(respons.Body)

		var d map[string]interface{}

		json.Unmarshal(responseBody, &d)

		fmt.Println("PUSH Respons Status = ", respons.Status)
		fmt.Println("PUSH Respons Body = ", d)
	}

	return nil
}
