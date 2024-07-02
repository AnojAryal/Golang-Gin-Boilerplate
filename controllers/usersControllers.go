package controllers

import (
	"net/http"

	"github.com/anojaryal/JWT-Authentication/initializers"
	"github.com/anojaryal/JWT-Authentication/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//get the details of req body
	var body struct {
		Email string
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return 
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	//create user
	user := models.User{Email: body.Email, Username: body.Username, Password:string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "failed to create user",
		})

		return
	}

	//respond
	c.JSON(http.StatusCreated, gin.H{})
}