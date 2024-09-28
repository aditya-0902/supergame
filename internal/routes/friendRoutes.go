package routes

import (
	"github.com/aditya-0902/supergaming/internal/controller"
	"github.com/gin-gonic/gin"
)

func FriendGroupRoutes(incomingRoute *gin.Engine) {

	friendsGroup := incomingRoute.Group("friends")
	{
		friendsGroup.POST("/request", controller.SendFriendRequestController())
		friendsGroup.POST("/accept", controller.AcceptFriendRequestController())
		friendsGroup.POST("/reject", controller.RejectFriendRequestController())
		friendsGroup.GET("/list", controller.GetFriendListController())
		friendsGroup.DELETE("/remove", controller.RemoveFriendController())
		friendsGroup.GET("/online-status", controller.FriendsOnlineStatusWebSocketController())
	}
}
