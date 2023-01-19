package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	ClickAction string `json:"click_action"`
	Amount      int    `json:"amount"`
}

func (h *Handler) whClient(c *gin.Context) {

	c.String(http.StatusOK, "OK")

	var input clientFromYTimes

	if err := c.BindJSON(&input); err != nil {

		logrus.Error("parsing error")
		return
	}

	userData, err := h.services.CoffeeDBUpdate.UpdatePoints(input.Phone, float32(input.PointsValue))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var outValue = int(userData.Value)

	err = pushRequest(outValue, userData.MessageKey)
	if err != nil {
		logrus.Error("pushRequest error")
	}
}

func (h *Handler) whMenu(c *gin.Context) {

	c.String(http.StatusOK, "OK")

	responseBody, err := getMenu()

	file, err := os.Create("response.json")
	if err != nil {
		logrus.Error("File create error")
	}
	defer file.Close()

	fmt.Fprintf(file, "%s", responseBody)

	data, err := h.services.CoffeeDBUpdate.UpdateDB()

	if err != nil {
		logrus.Error("Menu update error\n", err)
		return
	}

	logrus.Println("Successful menu update at ", data)
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
	req.Header.Add(authorizationHeader, os.Getenv("FBASE_KEY"))

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

func getMenu() ([]byte, error) {

	SHOP_GUID := os.Getenv("SHOP_GUID")
	URL := os.Getenv("URL") + "/menu/item/list?shopGuid=" + SHOP_GUID
	GUID := os.Getenv("GUID")

	req, err := http.NewRequest(
		"GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(authorizationHeader, GUID)

	respons, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer respons.Body.Close()

	responseBody, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
