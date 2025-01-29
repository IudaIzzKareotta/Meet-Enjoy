package handlers

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createEvent(c *gin.Context) {
	var event MeetEnjoy2.Event
	if err := c.BindJSON(&event); err != nil {
		logrus.Errorf("Error parsing json %s:", err)
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("Error getting user id %s:", err)
		return
	}

	id, err := h.services.Events.CreateEvent(event, userId)
	if err != nil {
		logrus.Errorf("Error creating event %s:", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getUserEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("Error getting user events %s:", err)
		return
	}

	events, err := h.services.GetUserEvents(userId)
	if err != nil {
		logrus.Errorf("Error creating event %s:", err)
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *Handler) getEventById(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("event_id"))
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

	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		logrus.Errorf("Error parsing eventId  %s:", err)
		return
	}

	if err := h.services.Events.UpdateEvent(eventId, updateInput); err != nil {
		logrus.Errorf("Error updating event %s:", err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		logrus.Errorf("Error parsing eventId %s:", err)
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("Error getting user id %s:", err)
		return
	}

	if err := h.services.Events.DeleteEvent(eventId, userId); err != nil {
		log.Println(err)
		logrus.Errorf("Error deleting event %s:", err)
		return
	}

	c.Status(http.StatusOK)
}
