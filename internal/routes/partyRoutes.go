package routes

import (
	"github.com/aditya-0902/supergaming/internal/controller"
	"github.com/gin-gonic/gin"
)

func PartyGroupRoutes(incomingRoute *gin.Engine) {

	partyGroup := incomingRoute.Group("party")
	{
		partyGroup.POST("/create", controller.CreateParty())
		partyGroup.POST("/invite", controller.InviteToParty())
		partyGroup.POST("/join", controller.JoinParty())
		partyGroup.DELETE("/leave", controller.LeaveParty())
		partyGroup.POST("/accept-invite", controller.AcceptInvite())
		partyGroup.POST("/reject-invite", controller.RejectInvite())
		partyGroup.DELETE("/user", controller.RemoveUserFromParty())
		partyGroup.GET("/party-status", controller.PartyStatusWebSocketController())
	}
}
