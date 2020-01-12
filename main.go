package main

import (
	"github.com/gin-gonic/gin"
	"social-network/controllers"
	"social-network/model"
)

func main() {
	router := gin.Default()
	controllers.SetRoutes(router)

	router.LoadHTMLGlob("templates/*")
	router.Static("/css", "static/css")
	router.Static("/img", "static/img")
	router.Static("/js", "static/js")

	model.ConnectDb("admin:admin@/social_network")
	defer model.CloseDbConnection()

	router.Run("0.0.0.0:8181") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
