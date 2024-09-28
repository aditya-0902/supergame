package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aditya-0902/supergaming/internal/global"
	"github.com/gorilla/websocket"
)

func CreateParty(leaderID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return partyService.CreateParty(ctx, leaderID)
}

func InviteToParty(partyID, inviterID, inviteeID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return partyService.InviteToParty(ctx, partyID, inviterID, inviteeID)
}

func JoinParty(leader, partyID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return partyService.JoinParty(ctx, leader, partyID, userID)
}

func LeaveParty(leaderID, partyID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return partyService.LeaveParty(ctx, leaderID, partyID, userID)
}

func AcceptInvite(leaderID, partyID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return partyService.AcceptInvite(ctx, leaderID, partyID, userID)
}

func RejectInvite(partyID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return partyService.RejectInvite(ctx, partyID, userID)
}

func RemoveUserFromParty(partyID, leaderID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	fmt.Println("party_id is ", partyID)
	party, err := partyService.GetPartyByID(ctx, partyID)
	if err != nil {
		fmt.Println("error occured", err.Error())
		return err
	}

	if party.PartyLeader != leaderID {
		return errors.New("only the leader can remove users from the party")
	}

	return partyService.RemoveUser(ctx, partyID, userID)
}

func (s *WebSocketService) PartyStatusWebSocketService(leaderId, partyID string, conn *websocket.Conn) {
	fmt.Println("leader and party are")
	fmt.Println(leaderId, partyID)
	s.RegisterPartyConnection(leaderId, partyID, conn)
	defer s.UnregisterPartyConnection(leaderId, partyID)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("WebSocket closed for leader %s: %v", leaderId, err)
			break
		}
	}
}

func (s *WebSocketService) RegisterPartyConnection(leaderId, partyID string, conn *websocket.Conn) {
	s.mu.Lock()
	s.connections[leaderId] = conn
	s.mu.Unlock()
	fmt.Printf("Registered connection for leader: %s, partyId %s connections: %+v\n", leaderId, partyID, s.connections)

	s.NotifyPartyStatus(leaderId, partyID)
}

func (s *WebSocketService) UnregisterPartyConnection(userID, partyID string) {
	s.mu.Lock()
	if conn, ok := s.connections[userID]; ok {
		conn.Close()
		delete(s.connections, userID)
	}
	s.mu.Unlock()

	s.NotifyPartyStatus(userID, partyID)
}
func (s *WebSocketService) NotifyPartyStatus(leaderId, partyID string) {
	s.mu.Lock()
	fmt.Println("Checking for party status map entry...")

	friends, ok := global.PartyStatusMap[leaderId][partyID]
	fmt.Printf("PartyStatusMap: %+v\n", global.PartyStatusMap)

	s.mu.Unlock()

	if !ok {
		return
	}

	statusMessage := map[string]interface{}{
		"type":     "joined members",
		"members":  friends,
		"party_id": partyID,
	}
	fmt.Println("hello   ", statusMessage)
	// for _, friendId := range friends {
	s.mu.Lock()
	fmt.Printf("Sending message to leader: %s for party: %s\n", leaderId, partyID)

	// fmt.Println(leaderId, partyID, friendId, friends)
	fmt.Printf("%+v", s.connections)
	fmt.Println("**************")
	if conn, ok := s.connections[leaderId]; ok {
		log.Printf("Sending message to leader %s", leaderId)
		if err := conn.WriteJSON(statusMessage); err != nil {
			log.Printf("Error sending message to leader %s: %v", leaderId, err)
		}
	} else {
		log.Printf("No connection found for leader %s", leaderId)
	}
	// if conn, ok := s.connections[friendId]; ok {
	// 	conn.WriteJSON(statusMessage)
	// }

	s.mu.Unlock()
	// }

}
