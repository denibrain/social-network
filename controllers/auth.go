package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"social-network/model"
	"strings"
)

const (
	userKey = "user"
)

func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
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

func getCurrentUser(c *gin.Context) string {
	session := sessions.Default(c)
	user := session.Get(userKey)
	if user == nil {
		return ""
	} else {
		return user.(string)
	}
}
