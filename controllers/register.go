package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/josephkipkemoi/politicapp/database"
	"github.com/josephkipkemoi/politicapp/utils"
)

type RegistrationInput struct {
	FirstName       string `validate:"required"`
	LastName        string `validate:"required"`
	PhoneNumber     int    `validate:"required"`
	Password        string `validate:"required"`
	ConfirmPassword string `validate:"required"`
}

func Register(c *gin.Context) {
	input := &RegistrationInput{}

	if err := c.ShouldBindJSON(input); err != nil {
		errs, ok := utils.ValidateRequest(err)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": errs,
			})
			return
		}
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
