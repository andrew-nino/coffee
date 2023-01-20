package handler

import (
	"coffee-app/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var client = &http.Client{}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth", h.appIdentity)
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.appIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllCategories)
		}

		items := api.Group("/items")
		{
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemFromCategory)
		}

		types := api.Group("/types")
		{
			types.GET("/:id", h.getTypesFromItem)
		}

		actions := api.Group("/actions")
		{
			actions.GET("/", h.getAllActions)
			actions.GET("/:id", h.userIdentity, h.getActionByID)
		}

		ballance := api.Group("/balance", h.userIdentity)
		{
			ballance.GET("/", h.getBalance)
		}
	}

	images := router.Group("/images")
	{
		images.GET("/:name", h.getImage)
	}

	//webhooks
	router.POST("/client/update", h.senderIdentity, h.whClient)

	router.POST("/menu/changed", h.senderIdentity, h.whMenu)

	return router
}
