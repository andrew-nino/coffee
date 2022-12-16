package handler

import (
	// "bytes"
	"coffee-app"
	// "encoding/json"
	// "fmt"
	"net/http"

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
	PhoneCode string `json:"phone_code" binding:"required" db:"phone_code"`
	Phone     string `json:"phone" binding:"required" db:"phone"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input sighInInput

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// reqjson, _ := json.Marshal(input)
	// fmt.Println(reqjson)
	// req, err := http.NewRequest(
	// 	"POST", "https://api.ytimes.ru/ex/client/loadClientInfo", bytes.NewBuffer(reqjson))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// // добавляем заголовок
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add(authorizationHeader, "f8946569-64a6-4772-b713-e9e001144576-1670230075741")

	// // Отправив на сервер, получаем ответ
	// client := &http.Client{}
	// respons, err := client.Do(req)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// defer respons.Body.Close()

	// decoder := json.NewDecoder(respons.Body)

	// var inputNew sighInInput

	// err = decoder.Decode(&inputNew)

	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	token, err := h.services.Authorization.GenerateToken(input.PhoneCode, input.Phone)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		// "data": inputNew,
	})
}
