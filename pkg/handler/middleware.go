package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader   = "Authorization"
	longTimeToken         = "Long-time-token"
	longTimeTokenPassword = "SuperPassword"
	authorizationHeaderDB = "AuthorizationDB"
	userCtx               = "userPhone"
	sendersUUID           = "24481d34-7498-11ed-a1eb-0242ac120002"
)

func (h *Handler) appIdentity(c *gin.Context) {

	ltt := c.GetHeader(longTimeToken)

	if ltt == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty ayth header")
		return
	}

	password, err := h.services.Authorization.ParseToken(ltt)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if password != longTimeTokenPassword {
		newErrorResponse(c, http.StatusUnauthorized, "password is not correct")
		return
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
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

	userPhone, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userPhone)
}

func (h *Handler) senderIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeaderDB)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty ayth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		logrus.Error("invalid auth header update-db")
		return
	}

	if headerParts[1] != sendersUUID {
		logrus.Error("invalid auth UUID")
		return
	}
}

func getUserId(c *gin.Context) (int, error) {

	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
		return 0, errors.New("user id is invalid type")
	}

	return idInt, nil
}
