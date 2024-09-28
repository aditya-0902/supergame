package utils

import (
	"fmt"
	"sync"

	"github.com/aditya-0902/supergaming/internal/global"
)

var friendsMutex sync.Mutex
var partyMutex sync.Mutex

func AddFriendToFriendOnlineStatusMap(user1, user2 string) {
	friendsMutex.Lock()
	fmt.Printf("%+v\n", global.FriendOnlineStatusMap)
	if _, exists := global.FriendOnlineStatusMap[user1]; !exists {
		global.FriendOnlineStatusMap[user1] = []string{}
	}
	global.FriendOnlineStatusMap[user1] = append(global.FriendOnlineStatusMap[user1], user2)

	if _, exists := global.FriendOnlineStatusMap[user2]; !exists {
		global.FriendOnlineStatusMap[user2] = []string{}
	}
	global.FriendOnlineStatusMap[user2] = append(global.FriendOnlineStatusMap[user2], user1)
	fmt.Printf("%+v", global.FriendOnlineStatusMap)
	fmt.Println()
	friendsMutex.Unlock()
}

func RemoveFriendFromFriendOnlineStatusMap(user1, user2 string) {
	friendsMutex.Lock()
	fmt.Printf("%+v", global.FriendOnlineStatusMap)
	if users, ok := global.FriendOnlineStatusMap[user1]; ok {
		var updatedUsers []string
		for _, user := range users {
			if user != user2 {
				updatedUsers = append(updatedUsers, user)
			}
		}

		global.FriendOnlineStatusMap[user1] = updatedUsers
	}

	if users, ok := global.FriendOnlineStatusMap[user2]; ok {
		var updatedUsers []string
		for _, user := range users {
			if user != user1 {
				updatedUsers = append(updatedUsers, user)
			}
		}

		global.FriendOnlineStatusMap[user2] = updatedUsers
	}
	fmt.Printf("%+v", global.FriendOnlineStatusMap)
	fmt.Println()
	friendsMutex.Unlock()
}

func AddParticipantToParty(leader, party, user string) {
	partyMutex.Lock()
	if _, exists := global.PartyStatusMap[leader]; !exists {
		global.PartyStatusMap[leader] = make(map[string][]string)
	}
	global.PartyStatusMap[leader][party] = append(global.PartyStatusMap[leader][party], user)
	partyMutex.Unlock()

}

func RemoveParticipantFromParty(leader, invitee, partyId string) {
	partyMutex.Lock()
	if invitees, ok := global.PartyStatusMap[leader][partyId]; ok {
		var updatedUsers []string
		for _, user := range invitees {
			if user != invitee {
				updatedUsers = append(updatedUsers, user)
			}
		}

		global.PartyStatusMap[leader][partyId] = updatedUsers
	}

	partyMutex.Unlock()
}
