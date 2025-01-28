package handlers

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createEvent(c *gin.Context) {
	var event MeetEnjoy2.Event
	if err := c.BindJSON(&event); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Error parsing event json")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error getting user id")
		return
	}

	id, err := h.services.Events.CreateEvent(event, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error creating event")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getUserEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error getting user id")
		return
	}

	events, err := h.services.GetUserEvents(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error getting user events")
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *Handler) getEventParticipants(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing eventId")
		return
	}

	participants, err := h.services.GetEventParticipants(eventId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Error getting user events")
		return
	}

	c.JSON(http.StatusOK, participants)
}

func (h *Handler) getEventById(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing eventId")
		return
	}

	event, err := h.services.Events.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error getting event")
		return
	}

	c.JSON(http.StatusOK, event)

}

func (h *Handler) updateEvent(c *gin.Context) {
	var updateInput MeetEnjoy2.UpdateEventInput
	if err := c.BindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing update event")
		return
	}

	eventId, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing eventId")
		return
	}

	if err := h.services.Events.UpdateEvent(eventId, updateInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Error updating event")
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing eventId")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error getting user id")
		return
	}

	if err := h.services.Events.DeleteEvent(eventId, userId); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Error deleting event")
		return
	}
}
