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

	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthenticationRequired())

	authGroup.GET("/", userHome)

	authGroup.GET("/user/:id", getSelectedUser)
	authGroup.GET("/user/:id/talks", getMessages)
	authGroup.POST("/user/:id/talks", sendMessage)
	authGroup.GET("/user/:id/feeds", getUserFeeds)
	authGroup.POST("/user/:id/feeds", postNewFeed)
	authGroup.POST("/user/:id/sendRequest", sendRequest)

	authGroup.GET("/user/:id/friends", getFriends)

	authGroup.POST("/my/friends", getFriends)
	authGroup.POST("/my/getRequest", sendRequest)
	authGroup.GET("/my/feeds", getFeeds)
	authGroup.POST("/my/feeds", postNewFeed)
	authGroup.POST("/searchUsers", searchUsers)

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
