package main

import (
	"meesh-server/endpoints"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	router.GET("/cmd", endpoints.GetCmd)
	router.POST("/cmd", endpoints.PostCmd)
	if _, ok := os.LookupEnv("MEESH_SERVER_USE_UNIX_SOCKET"); ok {
		router.RunUnix(os.Getenv("HOME") + "/public_html/meeshserver")
	} else {
		router.Run(":8080")
	}
}
