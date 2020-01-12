package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"social-network/model"
	"strconv"
	"strings"
)

const (
	userKey = "user"
)

func logIn(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("email")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	authed, err := model.AuthUser(username, password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if !authed {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userKey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	returnTo, ok := c.GetQuery("return_to")
	if !ok {
		returnTo = "/"
	}
	c.JSON(http.StatusOK, gin.H{"returnTo": returnTo})
}

func signUp(c *gin.Context) {
	newUser := make(map[string]string)
	var requiredFields = []string{"email", "password", "name", "surname", "age", "sex", "city"}
	for _, fieldName := range requiredFields {
		value := c.PostForm(fieldName)
		value = strings.Trim(value, " ")
		if value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field can't be empty", "field": fieldName})
			return
		}
		newUser[fieldName] = value
	}

	age, err := strconv.Atoi(newUser["age"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Age should be a number", "field": "age"})
		return
	}

	if age < 6 || age > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Age should be in range 6 to 100", "field": "age"})
		return
	}

	sex := newUser["sex"]
	if sex != "U" && sex != "F" && sex != "M" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value of sex", "field": "sex"})
		return
	}

	user := model.User{
		Login:     newUser["email"],
		Password:  newUser["password"],
		Name:      newUser["name"],
		Surname:   newUser["surname"],
		Age:       age,
		Sex:       sex,
		City:      newUser["city"],
		Interests: c.PostForm("interests"),
	}

	err = model.AddUser(user)
	if err != nil {
		if e, ok := err.(*model.DuplicateRecordError); ok {
			c.JSON(http.StatusConflict, gin.H{"error": e.Message, "field": "email"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	session := sessions.Default(c)
	// Save the username in the session
	session.Set(userKey, user.Login) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	returnTo, ok := c.GetQuery("return_to")
	if !ok {
		returnTo = "/"
	}
	c.JSON(http.StatusOK, gin.H{"returnTo": returnTo})
}

func signOut(c *gin.Context) {
	session := sessions.Default(c)
	// Save the username in the session
	session.Set(userKey, nil) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func getCurrentUser(c *gin.Context) string {
	session := sessions.Default(c)
	user := session.Get(userKey)
	if user == nil {
		return ""
	} else {
		return user.(string)
	}
}
