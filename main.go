package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"social-network/controllers"
	"social-network/model"
)

var (
	dsn      string
	rdsn     string
	endpoint string
)

func main() {
	initOptions()

	router := gin.Default()
	controllers.SetRoutes(router)

	router.LoadHTMLGlob("templates/*")
	router.Static("/ui", "static/ui")

	model.ConnectDb(dsn)
	defer model.CloseDb()

	if rdsn == "" {
		rdsn = dsn
	}

	model.ConnectReadonlyDb(rdsn)
	defer model.CloseReadonlyDb()

	router.Run(endpoint) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initOptions() {
	flag.StringVar(&dsn, "d", "admin:admin@/social_network", "DSN: user:pwd@host/database")
	flag.StringVar(&rdsn, "r", "", "Readonly DSN: user:pwd@host/database")
	flag.StringVar(&endpoint, "e", "0.0.0.0:8181", "Endpoint: 0.0.0.0:80")
	flag.Parse()
}
