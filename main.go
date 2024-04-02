package main

import (
	"log"

	"github.com/gin-gonic/gin"
	server "github.com/shash-786/HTTPgo/Server"
)

func main() {
	router := gin.Default()
	main_server := server.New()
	//DEFINE ROUTES
	router.POST("/user/create", main_server.CreateUser())
	router.GET("/user/search/:name", main_server.SearchUser())

	log.Fatal(router.Run(":8080"))
}
