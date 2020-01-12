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

	r.GET("/user/:id", getSelectedUser)

	r.GET("/", userList)

	// Debug functions
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})
}
