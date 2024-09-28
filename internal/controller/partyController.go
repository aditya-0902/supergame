package controller

import (
	"net/http"

	"github.com/aditya-0902/supergaming/internal/models"
	"github.com/aditya-0902/supergaming/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateParty() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req *models.CreatePartyRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		partyID, err := service.CreateParty(req.LeaderID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"party_id": partyID})
	}
}

func InviteToParty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.InvitePartyRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := service.InviteToParty(req.PartyID, req.InviterID, req.InviteeID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User invited to party"})
	}
}

func JoinParty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.JoinPartyRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := service.JoinParty(req.LeaderID, req.PartyID, req.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer PartySocketService.NotifyPartyStatus(req.LeaderID, req.PartyID)
		ctx.JSON(http.StatusOK, gin.H{"message": "User joined the party"})
	}
}

func LeaveParty() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req *models.LeavePartyRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := service.LeaveParty(req.LeaderID, req.PartyID, req.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer PartySocketService.NotifyPartyStatus(req.LeaderID, req.PartyID)
		ctx.JSON(http.StatusOK, gin.H{"message": "User left the party"})
	}
}

func AcceptInvite() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req *models.PartyInviteRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := service.AcceptInvite(req.LeaderID, req.PartyID, req.InviteeID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer PartySocketService.NotifyPartyStatus(req.LeaderID, req.PartyID)

		ctx.JSON(http.StatusOK, gin.H{"message": "Invitation accepted"})
	}
}
func RejectInvite() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req models.PartyInviteRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := service.RejectInvite(req.PartyID, req.InviteeID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Invitation rejected"})
	}
}

func RemoveUserFromParty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.RemoveUserFromPartyRequestModel
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := service.RemoveUserFromParty(req.PartyID, req.LeaderID, req.UserID); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer PartySocketService.NotifyPartyStatus(req.LeaderID, req.PartyID)

		ctx.JSON(http.StatusOK, gin.H{"message": "User removed from party"})
	}
}

func PartyStatusWebSocketController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to establish WebSocket connection"})
			return
		}
		defer conn.Close()

		leaderID := ctx.Query("leader_id")
		partyID := ctx.Query("party_id")

		PartySocketService.PartyStatusWebSocketService(leaderID, partyID, conn)
	}
}
