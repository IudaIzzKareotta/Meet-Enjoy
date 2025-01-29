package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getEventParticipants(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing eventId")
		return
	}

	participants, err := h.services.Participants.GetEventParticipants(eventId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Error getting user events")
		return
	}

	c.JSON(http.StatusOK, participants)
}

func (h *Handler) deleteParticipant(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("Error getting user id %s:", err)
	}

	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		logrus.Errorf("Error parsing event id %s:", err)
	}

	participantId, err := strconv.Atoi(c.Param("participant_id"))
	if err != nil {
		logrus.Errorf("Error parsing participant id %s:", err)
	}

	if err := h.services.Participants.DeleteParticipant(userId, eventId, participantId); err != nil {
		logrus.Errorf("Error deleting participant %s:", err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) createInvite(c *gin.Context) {
	//	eventId, err := strconv.Atoi(c.Param("event_id"))
	//	if err != nil {
	//		logrus.Errorf("Error parsing eventId %s:", err)
	//		return
	//	}
	//
	//	userId, err := getUserId(c)
	//	if err != nil {
	//		logrus.Errorf("Error getting user id %s:", err)
	//		return
	//	}
	//
	//	inviteToken := uuid.New().String()
	//	invite := MeetEnjoy2.Invite{
	//		EventId:   eventId,
	//		Token:     inviteToken,
	//		CreatedAt: time.Now(),
	//	}
	//
	//	fmt.Sprintf("http://localhost:8080/api/events/%s/participants/accept/%s", eventId, inviteToken)
}

func (h *Handler) acceptInvite(c *gin.Context) {

}
