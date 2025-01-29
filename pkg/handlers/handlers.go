package handlers

import (
	service2 "github.com/IudaIzzKareotta/Meet-Enjoy/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service2.Service
}

func NewHandler(services *service2.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(loggerMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp) // Регистрация
		auth.POST("sign-in", h.signIn) // Авторизация
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/events/:event_id/participants/accept/:inviteToken", h.acceptInvite)

		events := api.Group("/events")
		{
			events.GET("/", h.getUserEvents)                 // Получение всех мероприятий
			events.GET("/:event_id", h.getEventById)         // Получение одного мероприятия по id
			events.POST("/", h.createEvent)                  // Создание мероприятия
			events.POST("/:event_id/invite", h.createInvite) // Создание приглашения
			events.PATCH("/:event_id", h.updateEvent)        // Обновление мероприятия
			events.DELETE("/:event_id", h.deleteEvent)       // Удаление мероприятия\

			participants := events.Group("/:event_id/participants")
			{
				participants.GET("/", h.getEventParticipants)                // Получение учасников меропрития по id
				participants.PATCH("/:participant_id")                       // Обновление статуса участника
				participants.DELETE("/:participant_id", h.deleteParticipant) // Удаление участника из мероприятия
			}
		}
	}

	return router
}
