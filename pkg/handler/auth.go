package handler

import (
	"bytes"
	"coffee-app"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {

	var input coffee.User

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
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

type sighInInput struct {
	PhoneCode string `json:"phoneCode" binding:"required" db:"phone_code"`
	Phone     string `json:"phone" binding:"required" db:"phone"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input sighInInput

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	reqjson, _ := json.Marshal(input)

	URL := os.Getenv("URL")
	GUID := os.Getenv("GUID")

	req, err := http.NewRequest(
		"POST", URL, bytes.NewBuffer(reqjson))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(authorizationHeader, GUID)

	client := &http.Client{}
	respons, err := client.Do(req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer respons.Body.Close()

	responseBody, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var inputNew struct {
		Success bool `json:"success"`
		Rows    []struct {
			OrderPromoCount     float32 `json:"orderPromoCount"`
			OrderPromoFreeCount float32 `json:"orderPromoFreeCount"`
			Points              float32 `json:"points"`
		} `json:"rows"`
	}

	err = json.Unmarshal(responseBody, &inputNew)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var outDate float32

	if len(inputNew.Rows) != 0 && inputNew.Success != false {

		outDate, err = h.services.CoffeeDBUpdate.UpdatePoints(input.Phone, inputNew.Rows[0].Points)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	token, err := h.services.Authorization.GenerateToken(input.PhoneCode, input.Phone)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":  token,
		"points": outDate,
	})
}
