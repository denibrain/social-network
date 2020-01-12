package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-network/model"
	"strconv"
)

func getSelectedUser(c *gin.Context) {
	currentUser := getCurrentUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.html", nil)
		return
	}
	selectedUser, err := model.GetUser(id)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", nil)
		return
	}

	if selectedUser == nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.HTML(200, "user.html", gin.H{
		"currentUser": currentUser,
		"user":        selectedUser,
	})
}

func userList(c *gin.Context) {
	currentUser := getCurrentUser(c)
	users, err := model.GetUsers(20)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", nil)
		return
	}
	c.HTML(200, "index.html", gin.H{
		"title":       "Home",
		"users":       users,
		"currentUser": currentUser,
	})
}
