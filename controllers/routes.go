package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"social-network/middleware"
	"social-network/model"
	"strconv"
)

func noescape(str string) template.HTML {
	return template.HTML(str)
}

func SetRoutes(r *gin.Engine) {
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.SetFuncMap(template.FuncMap{
		"noescape": noescape,
	})

	authRequired := r.Group("/profile")
	authRequired.Use(middleware.AuthenticationRequired())
	authRequired.GET("/", func(c *gin.Context) {
		currentUser := getCurrentUser(c)
		c.HTML(200, "profile.html", gin.H{
			"currentUser": currentUser,
		})
	})

	r.GET("/signin", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title": "Sign in",
		})
	})

	r.POST("/signout", signOut)

	r.POST("/signin", logIn)

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", gin.H{
			"title": "Sign up",
		})
	})

	r.POST("/signup", signUp)

	r.GET("/user/:id", func(c *gin.Context) {
		currentUser := getCurrentUser(c)
		id, err := strconv.Atoi(c.Param("name"))
		if err != nil {
			c.HTML(http.StatusBadRequest, "400.html", nil)
			return
		}
		selectedUser, err := model.GetUser(id)
		if err != nil {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}
		c.JSON(200, gin.H{
			"message":      "pong",
			"currentUser":  currentUser,
			"selectedUser": selectedUser,
		})
	})

	r.GET("/", func(c *gin.Context) {
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
	})

	// Debug functions
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})
}
