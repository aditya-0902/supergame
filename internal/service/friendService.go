package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aditya-0902/supergaming/internal/global"
	"github.com/aditya-0902/supergaming/internal/models"
	"github.com/gorilla/websocket"
)

func SendFriendRequest(fromUser, toUser string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return friendService.AddFriendRequest(ctx, fromUser, toUser)

}

func AcceptFriendRequest(requestID, toUser string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return friendService.AcceptFriendRequest(ctx, toUser, requestID)
}

func RemoveFriend(userID, friendID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return friendService.RemoveFriend(ctx, userID, friendID)
}

func GetFriendsList(userID string) ([]models.Friend, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return friendService.GetFriendsList(ctx, userID)
}

func RejectFriendRequest(toUser, requestID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return friendService.RejectFriendRequest(ctx, toUser, requestID)
}

func (s *WebSocketService) FriendsOnlineStatusWebSocketService(userId string, conn *websocket.Conn) {
	s.RegisterFriendsConnection(userId, conn)
	defer s.UnregisterFriendsConnection(userId)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("WebSocket closed for user %s: %v", userId, err)
			break
		}
	}
}

func (s *WebSocketService) notifyFriendsOnlineStatus(userID string, online bool) {
	s.mu.Lock()
	friends, ok := global.FriendOnlineStatusMap[userID]
	s.mu.Unlock()

	if !ok {
		return
	}

	statusMessage := map[string]interface{}{
		"type":    "online_status",
		"user_id": userID,
		"online":  online,
	}
	fmt.Println("hello   ", statusMessage)

	for _, friendID := range friends {
		s.mu.Lock()
		fmt.Println(userID, friendID, friends)
		fmt.Printf("%+v", s.connections)
		fmt.Println("**************")
		if conn, ok := s.connections[friendID]; ok {
			log.Printf("Sending message to friend %s", friendID)
			if err := conn.WriteJSON(statusMessage); err != nil {
				log.Printf("Error sending message to friend %s: %v", friendID, err)
			}
		} else {
			log.Printf("No connection found for friend %s", friendID)
		}
		s.mu.Unlock()
	}
}

func (s *WebSocketService) RegisterFriendsConnection(userID string, conn *websocket.Conn) {
	s.mu.Lock()
	s.connections[userID] = conn
	s.mu.Unlock()
	// fmt.Printf("Registered connection for user: %s, connections: %+v\n", userID, s.connections)

	s.notifyFriendsOnlineStatus(userID, true)
}

func (s *WebSocketService) UnregisterFriendsConnection(userID string) {
	s.mu.Lock()
	if conn, ok := s.connections[userID]; ok {
		conn.Close()
		delete(s.connections, userID)
	}
	s.mu.Unlock()

	s.notifyFriendsOnlineStatus(userID, false)
}
