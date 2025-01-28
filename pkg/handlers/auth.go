package handlers

import (
	"MeetEnjoy"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user MeetEnjoy.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Error creating user")
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
		log.Print("err1")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		log.Print("err2")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
