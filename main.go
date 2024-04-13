package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	router2 "levi-apis/apis/router"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	router := gin.New()
	userRoute := router2.UserRouter{E: router}
	userRoute.Route()
	if err := router.Run(":1111"); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Server is running on port 1111")
	}
}
