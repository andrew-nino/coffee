package handler

import (
	"bytes"
	"coffee-app"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type sighInInput struct {
	PhoneCode string `json:"phoneCode" binding:"required" db:"phone_code"`
	Phone     string `json:"phone" binding:"required" db:"phone"`
}

type countsAndPoints struct {
	Success bool `json:"success"`
	Rows    []struct {
		OrderPromoCount     float32 `json:"orderPromoCount"`
		OrderPromoFreeCount float32 `json:"orderPromoFreeCount"`
		Points              float32 `json:"points"`
	}
}

func (h *Handler) signUp(c *gin.Context) {

	var input coffee.User
	var pointsReceived countsAndPoints

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var clientData = sighInInput{
		PhoneCode: input.PhoneCode,
		Phone:     input.Phone,
	}

	err := requestPointsInYTimes(clientData, &pointsReceived)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(pointsReceived.Rows) == 0 {

		err := createClientAndAddPointsInYTimes(input)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

	var input sighInInput
	var pointsReceived countsAndPoints
	var points float32

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.PhoneCode, input.Phone)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = requestPointsInYTimes(input, &pointsReceived)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(pointsReceived.Rows) != 0 {

		points, err = h.services.CoffeeDBUpdate.UpdatePoints(input.Phone, pointsReceived.Rows[0].Points)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":  token,
		"points": points,
	})
}

func requestPointsInYTimes(input sighInInput, pointsReceived *countsAndPoints) error {

	reqjson, _ := json.Marshal(input)

	URL := os.Getenv("URL") + "/client/loadClientInfo"
	GUID := os.Getenv("GUID")

	responseBody, err := QueryInYTimes(URL, GUID, reqjson)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, &pointsReceived); err != nil || pointsReceived.Success != true {
		return err
	}

	return nil
}

func createClientAndAddPointsInYTimes(input coffee.User) error {

	newUUID, err := exec.Command("uuidgen").Output()

	update := CliientUpdate{
		RequestId: string(newUUID),
		PhoneCode: input.PhoneCode,
		Phone:     input.Phone,
	}

	reqjson, _ := json.Marshal(update)

	URL := os.Getenv("URL") + "/client/createClientAndAddPoints"
	GUID := os.Getenv("GUID")

	responseBody, err := QueryInYTimes(URL, GUID, reqjson)

	if err != nil {
		return err
	}

	var counts_and_points countsAndPoints
	if err = json.Unmarshal(responseBody, &counts_and_points); err != nil || counts_and_points.Success != true {
		return err
	}

	return nil
}

func QueryInYTimes(url, guid string, reqjson []byte) ([]byte, error) {

	req, err := http.NewRequest(
		"POST", url, bytes.NewBuffer(reqjson))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(authorizationHeader, guid)

	client := &http.Client{}
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
