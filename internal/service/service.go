package service

import (
	"fmt"
	"sync"

	"github.com/aditya-0902/supergaming/internal/constants"
	"github.com/aditya-0902/supergaming/internal/database"
	"github.com/aditya-0902/supergaming/internal/database/mongod"
	"github.com/gorilla/websocket"
)

// var service database.FriendRepository
var (
	friendService database.FriendRepository
	partyService  database.PartyRepository
)

type WebSocketService struct {
	connections map[string]*websocket.Conn // userID -> WebSocket connection
	mu          sync.Mutex
}

func NewFriendWebSocketService() *WebSocketService {
	return &WebSocketService{
		connections: make(map[string]*websocket.Conn),
	}
}

func NewPartyWebSocketService() *WebSocketService {
	return &WebSocketService{
		connections: make(map[string]*websocket.Conn),
	}
}

func init() {
	var err error
	// service, err = database.NewMongoDatabase(constants.MONGO_URI)
	// if err != nil {
	// 	fmt.Printf("error occured %s", err)
	// }
	friendService, err = database.NewMongoDatabase(constants.MONGO_URI, mongod.WithCollection(constants.FRIEND_DATABASE, constants.FRIEND_COLLECTION))
	if err != nil {
		fmt.Printf("error occured %s", err)
	}
	partyService, err = database.NewMongoDatabase(constants.MONGO_URI, mongod.WithCollection(constants.PARTY_DATABASE, constants.PARTY_COLLECTION))
	if err != nil {
		fmt.Printf("error occured %s", err)
	}
}
