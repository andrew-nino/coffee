package handler

import (
	"coffee-app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getLists)
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

		update := api.Group("/update-db", h.senderIdentity)
		{
			update.PUT("/", h.updateDB)
		}
	}

	router.POST("/client/update", h.updateClient)

	return router
}
