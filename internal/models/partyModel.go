package models

import (
	"time"
)

type Party struct {
	PartyID     string       `bson:"party_id"`
	PartyLeader string       `bson:"party_leader"`
	Members     []string     `bson:"members"`
	Invitations []Invitation `bson:"invitations"`
	CreatedAt   time.Time    `bson:"created_at"`
}

type Invitation struct {
	InviteeID string    `bson:"invitee_id"`
	Status    string    `bson:"status"` // "pending", "accepted", "rejected"
	InvitedAt time.Time `bson:"invited_at"`
}
type CreatePartyRequestModel struct {
	LeaderID string `json:"leader_id"`
}
type InvitePartyRequestModel struct {
	PartyID   string `json:"party_id"`
	InviterID string `json:"inviter_id"`
	InviteeID string `json:"invitee_id"`
}
type JoinPartyRequestModel struct {
	LeaderID string `json:"leader_id"`
	PartyID  string `json:"party_id"`
	UserID   string `json:"user_id"`
}

type LeavePartyRequestModel struct {
	LeaderID string `json:"leader_id"`
	PartyID  string `json:"party_id"`
	UserID   string `json:"user_id"`
}

type PartyInviteRequestModel struct {
	LeaderID  string `json:"leader_id"`
	PartyID   string `json:"party_id"`
	InviteeID string `json:"user_id"`
}

type RemoveUserFromPartyRequestModel struct {
	PartyID  string `json:"party_id"`
	LeaderID string `json:"party_leader"`
	UserID   string `json:"user_id"`
}
