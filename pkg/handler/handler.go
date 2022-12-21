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

		// update := api.Group("/update-db", h.senderIdentity)
		// {
		// 	update.PUT("/", h.updateDB)
		// }
	}

	images := router.Group("/images", h.appIdentity)
	{
		go images.GET("/:name", h.getImage)
	}

	//webhooks
	router.POST("/client/update", h.updateClient)

	// router.POST("/menu/changed", h.updateMenu)

	return router
}
