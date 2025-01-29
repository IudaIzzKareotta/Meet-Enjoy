package handlers

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user MeetEnjoy2.User

	if err := c.BindJSON(&user); err != nil {
		logrus.Errorf("Error parsing json %s:", err)
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		logrus.Errorf("Error creating user %s:", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("Error parsing json %s:", err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		logrus.Errorf("Error generating token %s:", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
