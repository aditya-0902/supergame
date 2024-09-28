package controller

import "github.com/aditya-0902/supergaming/internal/service"

var (
	FriendSocketService = service.NewFriendWebSocketService()

	PartySocketService = service.NewPartyWebSocketService()
)
