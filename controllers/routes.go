package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"social-network/middleware"
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

	r.GET("/", middleware.AuthenticationRequired(), userHome)

	r.GET("/feeds", middleware.AuthenticationRequired(), getFeeds)
	r.POST("/feeds", middleware.AuthenticationRequired(), postNewFeed)

	r.GET("/user/:id", middleware.AuthenticationRequired(), getSelectedUser)
	r.GET("/user/:id/talks", middleware.AuthenticationRequired(), getMessages)
	r.POST("/user/:id/talks", middleware.AuthenticationRequired(), sendMessage)
	r.GET("/user/:id/feeds", middleware.AuthenticationRequired(), getUserFeeds)
	r.POST("/user/:id/sendRequest", middleware.AuthenticationRequired(), sendRequest)

	r.POST("/friends", middleware.AuthenticationRequired(), getFriends)
	r.POST("/searchUsers", middleware.AuthenticationRequired(), searchUsers)

	// Auth section
	r.POST("/signin", logIn)
	r.GET("/signin", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title": "Sign in",
			"css":   "signin",
		})
	})

	r.POST("/signout", signOut)

	r.POST("/signup", signUp)
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", gin.H{
			"title": "Sign up",
			"css":   "signup",
		})
	})

	// Debug functions
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})
}
