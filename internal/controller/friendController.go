package controller

import (
	"net/http"

	"github.com/aditya-0902/supergaming/internal/models"
	"github.com/aditya-0902/supergaming/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendFriendRequestController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request *models.FriendRequestModel
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := service.SendFriendRequest(request.FromUser, request.ToUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "accepted"})

	}
}

func AcceptFriendRequestController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.AcceptRequsetModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := service.AcceptFriendRequest(req.RequestID, req.ToUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Friend request accepted"})

	}
}

func RejectFriendRequestController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.AcceptRequsetModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := service.RejectFriendRequest(req.ToUser, req.RequestID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Friend request rejected"})

	}
}

func RemoveFriendController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.RejectRequestModel

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := service.RemoveFriend(req.UserID, req.FriendID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Friend removed"})
	}
}

func GetFriendListController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Query("user_id")
		if userID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
			return
		}

		friends, err := service.GetFriendsList(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"friends": friends})
	}
}

func FriendsOnlineStatusWebSocketController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to establish WebSocket connection"})
			return
		}
		defer conn.Close()

		userID := ctx.Query("user_id")

		FriendSocketService.FriendsOnlineStatusWebSocketService(userID, conn)

	}
}
