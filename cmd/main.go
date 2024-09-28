package main

import (
	"github.com/aditya-0902/supergaming/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.FriendGroupRoutes(router)
	routes.PartyGroupRoutes(router)
	router.Run(":8080")

}
