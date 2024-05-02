package http

import (
	"github.com/gin-gonic/gin"
	"to-do-checklist/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
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
			lists.POST("/", h.createList)
			lists.GET("/", h.getUsersLists)
			lists.GET("/:id", h.getList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := api.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_id", h.getItem)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}
	return router
}
