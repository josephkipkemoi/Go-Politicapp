package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/josephkipkemoi/politicapp/database"
)

type RegistrationInput struct {
	FirstName       string
	LastName        string
	PhoneNumber     int
	Password        string
	ConfirmPassword string
}

// var validate *validator.Validate

func Register(c *gin.Context) {
	input := &RegistrationInput{}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
	}
	// Validate user input
	ok, msg := validateInputPassword(input.Password, input.ConfirmPassword)
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": msg,
		})
	}

	u := database.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}

	user, err := u.AddUser()
	if err != nil {
		c.JSON(422, gin.H{
			"error": err,
		})
	}

	c.JSON(201, gin.H{
		"user": user,
	})
}

func validateInputPassword(p1, p2 string) (bool, string) {
	if !strings.EqualFold(p1, p2) {
		return false, "passwords do not match"
	}
	return true, "passwords match"
}
