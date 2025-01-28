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
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp) // Регистрация
		auth.POST("sign-in", h.signIn) // Авторизация
	}

	events := router.Group("/events", h.userIdentity)
	{
		events.POST("/", h.createEvent)           // Создание мероприятия
		events.GET("/", h.getUserEvents)          // Получение всех мероприятий
		events.GET("/:eventId", h.getEventById)   // Получение одного мероприятия по id
		events.PATCH("/:eventId", h.updateEvent)  // Обновление мероприятия
		events.DELETE("/:eventId", h.deleteEvent) // Удаление мероприятия
	}

	participants := router.Group("/participants", h.userIdentity)
	{
		participants.GET("/:eventId", h.getEventParticipants)
		participants.GET("/")
	}

	return router
}
